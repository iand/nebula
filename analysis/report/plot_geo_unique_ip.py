import pandas as pd
from lib_db import DBClient
from plot_geo import plot_geo

client = DBClient()
results = client.query(
    f"""
    WITH cte AS (
        SELECT v.peer_id, unnest(mas.multi_address_ids) multi_address_id
        FROM visits v
                 INNER JOIN multi_addresses_sets mas on mas.id = v.multi_addresses_set_id
        WHERE v.created_at > date_trunc('week', NOW() - '1 week'::interval)
          AND v.created_at < date_trunc('week', NOW())
        GROUP BY v.peer_id, unnest(mas.multi_address_ids)
    )
    SELECT ia.country, count(DISTINCT ia.address) count
    FROM multi_addresses ma
             INNER JOIN cte ON cte.multi_address_id = ma.id
             INNER JOIN multi_addresses_x_ip_addresses maxia on ma.id = maxia.multi_address_id
             INNER JOIN ip_addresses ia ON maxia.ip_address_id = ia.id
    GROUP BY ia.country
    ORDER BY count DESC
    """
)
results_df = pd.DataFrame(results, columns=["Country", "Count"])

plot_geo(results_df, "Unique IP Address", 800, "unique")
