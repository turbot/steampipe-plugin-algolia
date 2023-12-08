---
title: "Steampipe Table: algolia_log - Query Algolia Logs using SQL"
description: "Allows users to query Algolia Logs, providing insights into events and operations performed in the Algolia search platform."
---

# Table: algolia_log - Query Algolia Logs using SQL

Algolia is a powerful search-as-a-service solution, making it easy to build and optimize search experiences. It offers robust APIs and extensive documentation to customize the platform to specific needs. Algolia logs provide crucial insights into the operations and events occurring within the platform.

## Table Usage Guide

The `algolia_log` table provides insights into the logs within Algolia. As a developer or system administrator, explore log-specific details through this table, including the type of events, timestamps, and associated metadata. Utilize it to uncover information about operations, such as search queries, indexing operations, and API calls, providing a clear overview of activities within the Algolia platform.

## Examples

### List all log entries
Explore your entire log history in chronological order, giving you the ability to track changes, spot trends and identify potential issues over time. This is particularly useful for auditing, debugging, and maintaining the overall health of your system.

```sql+postgres
select
  *
from
  algolia_log
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  algolia_log
order by
  timestamp desc;
```

### List errors from the log
Identify instances where requests to Algolia have resulted in errors. This can be used to pinpoint specific issues and improve the overall performance of your application.

```sql+postgres
select
  *
from
  algolia_log
where
  answer_code != 200
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  algolia_log
where
  answer_code != 200
order by
  timestamp desc;
```

### Find slow queries (>5ms) in the last 24 hours
Explore which queries have been performing slowly, taking more than 5 milliseconds, in the last 24 hours. This can be useful for identifying potential bottlenecks and improving overall system performance.

```sql+postgres
select
  *
from
  algolia_log
where
  processing_time_ms > 5
  and age(timestamp) < interval '24 hours'
order by
  timestamp desc;
```

```sql+sqlite
select
  *
from
  algolia_log
where
  processing_time_ms > 5
  and julianday('now') - julianday(timestamp) < 1
order by
  timestamp desc;
```

### Log entries by user agent
Explore the frequency of different user agents accessing your system to identify trends or anomalies. This can assist in understanding user behaviors, detecting potential security risks, and optimizing system performance.

```sql+postgres
select
  query_headers ->> 'User-Agent' as user_agent,
  count(*)
from
  algolia_log
group by
  user_agent
order by
  count desc;
```

```sql+sqlite
select
  json_extract(query_headers, '$."User-Agent"') as user_agent,
  count(*)
from
  algolia_log
group by
  user_agent
order by
  count(*) desc;
```