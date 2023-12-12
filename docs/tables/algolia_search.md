---
title: "Steampipe Table: algolia_search - Query Algolia Search Indices using SQL"
description: "Allows users to query Algolia Search Indices, specifically the search operations performed, providing insights into search patterns and potential optimizations."
---

# Table: algolia_search - Query Algolia Search Indices using SQL

Algolia is a hosted search engine API that provides developers with a powerful and flexible search solution. It offers a variety of features such as typo-tolerance, filters and facets, geo search, and language-agnostic full-text search. Algolia's Search Indices are the core resources where search operations are performed, and they contain the searchable data of an application.

## Table Usage Guide

The `algolia_search` table provides insights into search operations performed within Algolia's Search Indices. As a developer or data analyst, explore search-specific details through this table, including query text, number of hits, and processing time. Utilize it to uncover information about search patterns, such as frequently searched queries or queries with high processing time, aiding in potential optimizations and improvements to your application's search functionality.

## Examples

### Search the offices index for the term "usa"
Analyze the 'offices' index to pinpoint specific entries related to 'USA'. This can be beneficial in understanding the distribution and presence of offices across the United States.

```sql+postgres
select
  rank,
  object_id,
  jsonb_pretty(hit) as result
from
  algolia_search
where
  index = 'offices'
  and query = 'usa'
order by
  rank;
```

```sql+sqlite
select
  rank,
  object_id,
  hit as result
from
  algolia_search
where
  index = 'offices'
  and query = 'usa'
order by
  rank;
```

### Get all records from all indexes
Explore all indexed records to gain insights into the data stored in your Algolia search system. This can be useful for understanding the breadth and variety of information available for search operations.
Warning: This is a very large query, use carefully!


```sql+postgres
select
  *
from
  algolia_search;
```

```sql+sqlite
select
  *
from
  algolia_search;
```