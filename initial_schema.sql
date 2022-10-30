DROP TABLE stocks_real_time;
CREATE TABLE stocks_real_time (
  time timestamp without time zone NOT NULL,
  symbol TEXT NOT NULL,
  price DOUBLE PRECISION NULL,
  day_volume INT NULL
);

SELECT create_hypertable('stocks_real_time', 'time', create_default_indexes=>FALSE, chunk_time_interval=>INTERVAL '1 day');
CREATE INDEX ix_symbol_time ON stocks_real_time (symbol, time DESC);


-- 1min k-line
DROP MATERIALIZED VIEW stock_candlestick_1m;
CREATE MATERIALIZED VIEW stock_candlestick_1m
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 min', "time") AS bucket,
  symbol,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM stocks_real_time srt
GROUP BY bucket, symbol
WITH NO DATA;




SELECT set_chunk_time_interval('_timescaledb_internal._materialized_hypertable_9', INTERVAL '3000 mins');

SELECT add_continuous_aggregate_policy('stock_candlestick_1m',
  start_offset => INTERVAL '7 days',
  end_offset => INTERVAL '1 m',
  schedule_interval => INTERVAL '1 m');

-- compression
ALTER MATERIALIZED VIEW stock_candlestick_1m set (timescaledb.compress = true);
SELECT add_compression_policy('stock_candlestick_1m', compress_after=>'14 days'::interval);


  -- 1hour k-line
DROP MATERIALIZED VIEW stock_candlestick_1h;
CREATE MATERIALIZED VIEW stock_candlestick_1h
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 hour', "time") AS bucket,
  symbol,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM stocks_real_time srt
GROUP BY bucket, symbol
WITH NO DATA;

SELECT add_continuous_aggregate_policy('stock_candlestick_1h',
  start_offset => INTERVAL '1 month',
  end_offset => INTERVAL '1 h',
  schedule_interval => INTERVAL '1 h');


-- 1day k-line
DROP MATERIALIZED VIEW stock_candlestick_1d;
CREATE MATERIALIZED VIEW stock_candlestick_1d
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 day', "time") AS bucket,
  symbol,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM stocks_real_time srt
GROUP BY bucket, symbol
WITH NO DATA;

SELECT add_continuous_aggregate_policy('stock_candlestick_1d',
  start_offset => INTERVAL '1 month',
  end_offset => INTERVAL '1 d',
  schedule_interval => INTERVAL '1 d');



-- query 30 min k-line from 1min table
SELECT
  time_bucket('30 min', "bucket") AS time, symbol, max(high), first(open, "bucket"), last(close, "bucket"), min(low)
FROM stock_candlestick_1m
where symbol = 'AAPL'
GROUP BY time, symbol
limit 100