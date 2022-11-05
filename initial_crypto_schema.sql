CREATE TABLE public.trade (
	"time" timestamp NOT NULL DEFAULT now(),
	order_id varchar NOT NULL DEFAULT ''::character varying,
	market varchar NOT NULL DEFAULT ''::character varying,
	side int2 NOT NULL DEFAULT 0,
	price numeric(20, 8) NOT NULL DEFAULT 0,
	"size" numeric(20, 8) NOT NULL DEFAULT 0
);

CREATE INDEX trade_market_idx ON public.trade USING btree (market, "time" DESC);
SELECT create_hypertable('public.trade', 'time', create_default_indexes=>FALSE, chunk_time_interval=>INTERVAL '1 day');


-- 1 min k-line
CREATE MATERIALIZED VIEW k_line_1m
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 min', "time") AS bucket,
  market,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM trade
GROUP BY bucket, market
WITH NO DATA;

SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_2', INTERVAL '3000 mins');

SELECT add_continuous_aggregate_policy('k_line_1m',
  start_offset => INTERVAL '1 days',
  end_offset => INTERVAL '1 m',
  schedule_interval => INTERVAL '1 m');



-- 1hour k-line
CREATE MATERIALIZED VIEW k_line_1h
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 hour', "time") AS bucket,
  market,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM trade
GROUP BY bucket, market
WITH NO DATA;

SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_3', INTERVAL '800 hours');

SELECT add_continuous_aggregate_policy('k_line_1h',
  start_offset => INTERVAL '2 days',
  end_offset => INTERVAL '1 h',
  schedule_interval => INTERVAL '1 h');


-- 1day k-line
CREATE MATERIALIZED VIEW k_line_1d
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 day', "time") AS bucket,
  market,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM trade
GROUP BY bucket, market
WITH NO DATA;

SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_4', INTERVAL '3000 day');

SELECT add_continuous_aggregate_policy('k_line_1d',
  start_offset => INTERVAL '3 days',
  end_offset => INTERVAL '1 d',
  schedule_interval => INTERVAL '1 d');