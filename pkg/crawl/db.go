package crawl

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"go.opencensus.io/stats"

	"github.com/dennis-tra/nebula-crawler/pkg/metrics"
	"github.com/dennis-tra/nebula-crawler/pkg/models"
)

func (s *Scheduler) queryPeers(pis []peer.AddrInfo) error {
	var queryList []peer.AddrInfo
	for _, pi := range pis {
		if _, crawled := s.crawled[pi.ID]; crawled {
			continue
		}
		if _, loaded := s.dbPeers[pi.ID]; loaded {
			continue
		}
		queryList = append(queryList, pi)
	}

	if len(queryList) == 0 {
		return nil
	}

	peers, err := s.dbc.QueryPeers(s.ServiceContext(), queryList)
	if err != nil {
		return err
	}

	for _, p := range peers {
		mh, err := peer.Decode(p.MultiHash)
		if err != nil {
			// TODO: log
			continue
		}
		s.dbPeers[mh] = p
	}

	return nil
}

// persistCrawlResult inserts the given peer with its multi addresses in the database and
// upserts its currently active session
func (s *Scheduler) persistCrawlResult(cr Result) error {
	var err error
	ctx := s.ServiceContext()
	dbPeer, found := s.dbPeers[cr.Peer.ID]
	if !found || dbPeer.AgentVersion.String != cr.Agent {
		if dbPeer, err = s.dbc.UpsertPeer(ctx, cr.Peer, cr.Agent); err != nil {
			return errors.Wrap(err, "upsert peer")
		}
	}

	startUpsert := time.Now()
	if cr.Error == nil {
		if err := s.dbc.UpsertSessionSuccess(dbPeer); err != nil {
			return errors.Wrap(err, "upsert session success")
		}
	} else if cr.Error != s.ServiceContext().Err() {
		if err := s.dbc.UpsertSessionError(dbPeer, cr.ErrorTime, determineDialError(cr.Error)); err != nil {
			return errors.Wrap(err, "upsert session error")
		}
	}
	stats.Record(s.ServiceContext(), metrics.CrawledUpsertDuration.M(millisSince(startUpsert)))

	// Persist latency measurements
	if cr.Latencies != nil {
		if err := s.dbc.InsertLatencies(ctx, dbPeer, cr.Latencies); err != nil {
			return errors.Wrap(err, "insert latencies")
		}
	}

	return nil
}

// updateCrawl writes crawl statistics to the database. TODO: comment
func (s *Scheduler) updateCrawl(ctx context.Context) error {
	log.Infoln("Persisting crawl result...")

	s.crawl.StartedAt = s.StartTime
	s.crawl.FinishedAt = null.TimeFrom(time.Now())
	s.crawl.CrawledPeers = null.IntFrom(len(s.crawled))
	s.crawl.DialablePeers = null.IntFrom(len(s.crawled) - s.TotalErrors())
	s.crawl.UndialablePeers = null.IntFrom(s.TotalErrors())

	if s.ServiceContext().Err() == nil {
		s.crawl.State = models.CrawlStateSucceeded
	} else if errors.Is(s.ServiceContext().Err(), context.Canceled) {
		s.crawl.State = models.CrawlStateCancelled
	} else {
		s.crawl.State = models.CrawlStateFailed
	}

	return s.dbc.UpdateCrawl(ctx, s.crawl)
}

// persistPeerProperties writes peer property statistics to the database.
func (s *Scheduler) persistPeerProperties(ctx context.Context) error {
	log.Infoln("Persisting peer properties...")

	// Extract full and core agent versions. Core agent versions are just strings like 0.8.0 or 0.5.0
	// The full agent versions have much more information e.g., /go-ipfs/0.4.21-dev/789dab3
	avFull := map[string]int{}
	avCore := map[string]int{}
	for version, count := range s.agentVersion {
		avFull[version] += count
		matches := agentVersionRegex.FindStringSubmatch(version)
		if matches != nil {
			avCore[matches[1]] += count
		}
	}
	pps := map[string]map[string]int{
		"agent_version":      avFull,
		"agent_version_core": avCore,
		"protocol":           s.protocols,
		"error":              s.errors,
	}

	return s.dbc.PersistPeerProperties(ctx, s.crawl, pps)
}