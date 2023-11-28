# learning-timescale

1. 目前 timescaledb 无法做到 Nested Continuous Aggregates
raw data → 1m → 1h → 1d
<https://github.com/timescale/timescaledb/pull/4668>

1. 如果想調整 Continuous Aggregates 的 chunk sizing
<https://docs.timescale.com/timescaledb/latest/how-to-guides/hypertables/change-chunk-intervals/>
<https://stackoverflow.com/questions/67833467/set-chunk-time-interval-on-the-continuous-aggregates-materialization-view>

1. Continuous Aggregates 支持整點或固定的觸發時間背景任務
<https://github.com/timescale/timescaledb/pull/4664>

1. Check chunks size
SELECT * FROM chunks_detailed_size('dist_table')

1. 自然月用　month,　不是 30d
<https://docs.timescale.com/api/latest/hyperfunctions/time_bucket/>

## Compression

1. Video
<https://www.youtube.com/watch?v=RR1xayRusBc>

1. data can't be changed after compression
<https://docs.timescale.com/timescaledb/latest/how-to-guides/compression/modify-a-schema/>

## ANALYZE

EXPLAIN (ANALYZE,BUFFERS)
select *
from stock_candlestick_1m scm
where symbol = 'AAPL'
order by bucket desc
limit 1

## database mangagement

1. 列出 hypertable

```sql
SELECT * FROM timescaledb_information.hypertables
```
