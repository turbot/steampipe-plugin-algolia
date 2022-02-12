# Table: algolia_log

Recent log entries across all indexes and operation types. Results are limited to last 1000 entries and/or a maximum of 7 days age.

## Examples

### List all log entries

```sql
select
  *
from
  algolia_log
order by
  timestamp desc
```

### List errors from the log

```sql
select
  *
from
  algolia_log
where
  answer_code != 200
order by
  timestamp desc
```

### Find slow queries (>5ms) in the last 24 hours

```sql
select
  *
from
  algolia_log
where
  processing_time_ms > 5
  and age(timestamp) < interval '24 hours'
order by
  timestamp desc
```

### Log entries by user agent

```sql
select
  query_headers ->> 'User-Agent' as user_agent,
  count(*)
from
  algolia_log
group by
  user_agent
order by
  count desc
```
