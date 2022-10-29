CREATE TABLE stocks_real_time (
  time timestamp without time zone NOT NULL,
  symbol TEXT NOT NULL,
  price DOUBLE PRECISION NULL,
  day_volume INT NULL
);

CREATE INDEX ix_symbol_time ON stocks_real_time (symbol, time DESC);
SELECT create_hypertable('stocks_real_time', 'time', chunk_time_interval=>INTERVAL '1100 seconds');


CREATE MATERIALIZED VIEW stock_candlestick_1m
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 min', "time") AS min,
  symbol,
  max(price) AS high,
  first(price, time) AS open,
  last(price, time) AS close,
  min(price) AS low
FROM stocks_real_time srt
GROUP BY min, symbol;

