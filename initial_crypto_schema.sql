CREATE TABLE public.trade (
    "time" timestamp NOT NULL DEFAULT now(),
	order_id varchar NOT NULL DEFAULT ''::character varying,
	market varchar NOT NULL DEFAULT ''::character varying,
    side int2 NOT NULL DEFAULT 0,
	price numeric NOT NULL DEFAULT 0,
	"size" numeric NOT NULL DEFAULT 0,
	vol numeric NOT NULL DEFAULT 0
);

CREATE INDEX trade_market_idx ON public.trade USING btree (market, "time" DESC);
SELECT create_hypertable('public.trade', 'time', create_default_indexes=>FALSE, chunk_time_interval=>INTERVAL '1 day');
