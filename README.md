# learning-timescale

1. 目前 timescaledb 无法做到 Nested Continuous Aggregates
raw data → 1m → 1h → 1d
<https://github.com/timescale/timescaledb/pull/4668>

1. 如果想調整Continuous Aggregates 的 chunk sizing
<https://stackoverflow.com/questions/67833467/set-chunk-time-interval-on-the-continuous-aggregates-materialization-view>

1. Continuous Aggregates 目前不支持整點觸發, 必須使用 cronjob
