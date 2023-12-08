---
title: "Steampipe Table: algolia_index - Query Algolia Indices using SQL"
description: "Allows users to query Algolia Indices, providing insights into index configurations and associated data."
---

# Table: algolia_index - Query Algolia Indices using SQL

Algolia is a powerful search-as-a-service solution, making it easy to build and manage search functionality for your websites and mobile applications. Indices in Algolia represent the data sets that are searchable by a search interface. They contain a collection of records, each one representing a specific object, and can be configured with various settings to fine-tune how the data is processed, stored, and retrieved.

## Table Usage Guide

The `algolia_index` table provides insights into Indices within Algolia. As a developer or data analyst, explore index-specific details through this table, including settings, task status, and associated metadata. Utilize it to uncover information about indices, such as their configuration, the status of indexing tasks, and the retrieval of index-specific data.

## Examples

### List all indices
Discover the segments that have been created within a certain timeframe to understand the volume of data entries. This can help in assessing the performance and efficiency of data indexing within your system.

```sql+postgres
select
  name,
  created_at,
  entries
from
  algolia_index
order by
  name;
```

```sql+sqlite
select
  name,
  created_at,
  entries
from
  algolia_index
order by
  name;
```

### Top 5 largest indices
Discover the segments that have the highest number of entries in your Algolia index. This allows you to identify areas of your database that are heavily populated and might require optimization or closer review.

```sql+postgres
select
  name,
  entries,
  data_size
from
  algolia_index
order by
  entries desc;
```

```sql+sqlite
select
  name,
  entries,
  data_size
from
  algolia_index
order by
  entries desc;
```

### Indices last updated more than 7 days ago
Discover the segments that have not been updated in the last 7 days. This can be useful in identifying stale or outdated data that may need refreshing or removal.

```sql+postgres
select
  name,
  updated_at
from
  algolia_index
where
  age(updated_at) > interval '7 days';
```

```sql+sqlite
select
  name,
  updated_at
from
  algolia_index
where
  julianday('now') - julianday(updated_at) > 7;
```