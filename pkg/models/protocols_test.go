// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProtocols(t *testing.T) {
	t.Parallel()

	query := Protocols()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProtocolsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProtocolsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Protocols().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProtocolsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProtocolSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProtocolsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProtocolExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Protocol exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProtocolExists to return true, but got false.")
	}
}

func testProtocolsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	protocolFound, err := FindProtocol(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if protocolFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProtocolsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Protocols().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProtocolsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Protocols().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProtocolsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	protocolOne := &Protocol{}
	protocolTwo := &Protocol{}
	if err = randomize.Struct(seed, protocolOne, protocolDBTypes, false, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}
	if err = randomize.Struct(seed, protocolTwo, protocolDBTypes, false, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = protocolOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = protocolTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Protocols().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProtocolsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	protocolOne := &Protocol{}
	protocolTwo := &Protocol{}
	if err = randomize.Struct(seed, protocolOne, protocolDBTypes, false, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}
	if err = randomize.Struct(seed, protocolTwo, protocolDBTypes, false, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = protocolOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = protocolTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func protocolBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func protocolAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Protocol) error {
	*o = Protocol{}
	return nil
}

func testProtocolsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Protocol{}
	o := &Protocol{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, protocolDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Protocol object: %s", err)
	}

	AddProtocolHook(boil.BeforeInsertHook, protocolBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	protocolBeforeInsertHooks = []ProtocolHook{}

	AddProtocolHook(boil.AfterInsertHook, protocolAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	protocolAfterInsertHooks = []ProtocolHook{}

	AddProtocolHook(boil.AfterSelectHook, protocolAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	protocolAfterSelectHooks = []ProtocolHook{}

	AddProtocolHook(boil.BeforeUpdateHook, protocolBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	protocolBeforeUpdateHooks = []ProtocolHook{}

	AddProtocolHook(boil.AfterUpdateHook, protocolAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	protocolAfterUpdateHooks = []ProtocolHook{}

	AddProtocolHook(boil.BeforeDeleteHook, protocolBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	protocolBeforeDeleteHooks = []ProtocolHook{}

	AddProtocolHook(boil.AfterDeleteHook, protocolAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	protocolAfterDeleteHooks = []ProtocolHook{}

	AddProtocolHook(boil.BeforeUpsertHook, protocolBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	protocolBeforeUpsertHooks = []ProtocolHook{}

	AddProtocolHook(boil.AfterUpsertHook, protocolAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	protocolAfterUpsertHooks = []ProtocolHook{}
}

func testProtocolsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProtocolsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(protocolColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProtocolToManyCrawlProperties(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Protocol
	var b, c CrawlProperty

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, crawlPropertyDBTypes, false, crawlPropertyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, crawlPropertyDBTypes, false, crawlPropertyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ProtocolID, a.ID)
	queries.Assign(&c.ProtocolID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CrawlProperties().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ProtocolID, b.ProtocolID) {
			bFound = true
		}
		if queries.Equal(v.ProtocolID, c.ProtocolID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProtocolSlice{&a}
	if err = a.L.LoadCrawlProperties(ctx, tx, false, (*[]*Protocol)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CrawlProperties); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CrawlProperties = nil
	if err = a.L.LoadCrawlProperties(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CrawlProperties); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProtocolToManyAddOpCrawlProperties(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Protocol
	var b, c, d, e CrawlProperty

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, protocolDBTypes, false, strmangle.SetComplement(protocolPrimaryKeyColumns, protocolColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CrawlProperty{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, crawlPropertyDBTypes, false, strmangle.SetComplement(crawlPropertyPrimaryKeyColumns, crawlPropertyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*CrawlProperty{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCrawlProperties(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ProtocolID) {
			t.Error("foreign key was wrong value", a.ID, first.ProtocolID)
		}
		if !queries.Equal(a.ID, second.ProtocolID) {
			t.Error("foreign key was wrong value", a.ID, second.ProtocolID)
		}

		if first.R.Protocol != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Protocol != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CrawlProperties[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CrawlProperties[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CrawlProperties().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testProtocolToManySetOpCrawlProperties(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Protocol
	var b, c, d, e CrawlProperty

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, protocolDBTypes, false, strmangle.SetComplement(protocolPrimaryKeyColumns, protocolColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CrawlProperty{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, crawlPropertyDBTypes, false, strmangle.SetComplement(crawlPropertyPrimaryKeyColumns, crawlPropertyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetCrawlProperties(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CrawlProperties().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCrawlProperties(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CrawlProperties().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ProtocolID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ProtocolID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ProtocolID) {
		t.Error("foreign key was wrong value", a.ID, d.ProtocolID)
	}
	if !queries.Equal(a.ID, e.ProtocolID) {
		t.Error("foreign key was wrong value", a.ID, e.ProtocolID)
	}

	if b.R.Protocol != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Protocol != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Protocol != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Protocol != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.CrawlProperties[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.CrawlProperties[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testProtocolToManyRemoveOpCrawlProperties(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Protocol
	var b, c, d, e CrawlProperty

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, protocolDBTypes, false, strmangle.SetComplement(protocolPrimaryKeyColumns, protocolColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CrawlProperty{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, crawlPropertyDBTypes, false, strmangle.SetComplement(crawlPropertyPrimaryKeyColumns, crawlPropertyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCrawlProperties(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CrawlProperties().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCrawlProperties(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CrawlProperties().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ProtocolID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ProtocolID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Protocol != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Protocol != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Protocol != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Protocol != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.CrawlProperties) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.CrawlProperties[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.CrawlProperties[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testProtocolsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProtocolsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProtocolSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProtocolsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Protocols().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	protocolDBTypes = map[string]string{`ID`: `integer`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`, `Protocol`: `character varying`}
	_               = bytes.MinRead
)

func testProtocolsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(protocolPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(protocolAllColumns) == len(protocolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProtocolsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(protocolAllColumns) == len(protocolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Protocol{}
	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, protocolDBTypes, true, protocolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(protocolAllColumns, protocolPrimaryKeyColumns) {
		fields = protocolAllColumns
	} else {
		fields = strmangle.SetComplement(
			protocolAllColumns,
			protocolPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProtocolSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProtocolsUpsert(t *testing.T) {
	t.Parallel()

	if len(protocolAllColumns) == len(protocolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Protocol{}
	if err = randomize.Struct(seed, &o, protocolDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Protocol: %s", err)
	}

	count, err := Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, protocolDBTypes, false, protocolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Protocol struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Protocol: %s", err)
	}

	count, err = Protocols().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}