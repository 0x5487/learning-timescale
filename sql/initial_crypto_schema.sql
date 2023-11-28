CREATE TABLE public.trade (
	"time" timestamp NOT NULL DEFAULT now(),
  market varchar NOT NULL DEFAULT ''::character varying,
	id varchar NOT NULL DEFAULT ''::character varying,
	side int2 NOT NULL DEFAULT 0,
	price numeric(20, 8) NOT NULL DEFAULT 0,
	"size" numeric(20, 8) NOT NULL DEFAULT 0,
  "volume" numeric(20, 8) NOT NULL DEFAULT 0,
  CONSTRAINT pk_trade PRIMARY KEY ("time", id)
);
CREATE INDEX idx_trade_market ON public.trade USING btree ("time" DESC, market);
SELECT create_hypertable('public.trade', 'time', partitioning_column => 'market', number_partitions => 2, chunk_time_interval => interval '1 hour',create_default_indexes=>FALSE);



-- 1 min k-line
-- SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_2', INTERVAL '3000 mins');
CREATE MATERIALIZED VIEW k_line_1m
WITH (timescaledb.continuous,
timescaledb.materialized_only = true,
timescaledb.chunk_time_interval = 3000 * 60 * 1000000) AS
SELECT
  time_bucket('1 min', "time") AS bucket,
  market,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low,
  sum(size) as vol_base,
  sum(volume) as vol_quote
FROM public.trade
GROUP BY bucket, market
ORDER BY bucket
WITH NO DATA;

SELECT add_continuous_aggregate_policy('k_line_1m',
  start_offset => INTERVAL '5 m',
  end_offset => NULL,
  schedule_interval => INTERVAL '1 s');



-- 1hour k-line
-- SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_3', INTERVAL '800 hours');
CREATE MATERIALIZED VIEW k_line_1h
WITH (timescaledb.continuous,
timescaledb.materialized_only = true) AS
SELECT
  time_bucket('1 hour', "bucket") AS bucket,
  market,
  max(high) AS high,
  first(open, bucket) AS open,
  last(close, bucket) AS close,
  min(low) AS low,
  sum(vol_base) as vol_base,
  sum(vol_quote) as vol_quote
FROM k_line_1m
GROUP BY 1, 2
WITH NO DATA;


SELECT add_continuous_aggregate_policy('k_line_1h',
  start_offset => INTERVAL '12 h',
  end_offset => NULL,
  schedule_interval => INTERVAL '1 s');


-- 1day k-line
-- SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_4', INTERVAL '3000 day');
CREATE MATERIALIZED VIEW k_line_1d
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 day', "bucket") AS bucket,
  market,
  max(high) AS high,
  first(open, bucket) AS open,
  last(close, bucket) AS close,
  min(low) AS low,
  sum(vol_base) as vol_base,
  sum(vol_quote) as vol_quote
FROM k_line_1h
GROUP BY 1,2
WITH NO DATA;


SELECT add_continuous_aggregate_policy('k_line_1d',
  start_offset => INTERVAL '1 days',
  end_offset => NULL,
  schedule_interval => INTERVAL '1 s');