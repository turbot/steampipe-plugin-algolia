---
title: "Steampipe Table: algolia_search_metadata - Query Algolia Search Metadata using SQL"
description: "Allows users to query Algolia Search Metadata, specifically providing insights into search operations and their associated metadata."
---

# Table: algolia_search_metadata - Query Algolia Search Metadata using SQL

Algolia is a search and discovery API platform that provides a seamless search and discovery experience across devices and platforms. It enables developers to build search functionality into their applications, websites, and software. Algolia's Search Metadata is a resource that provides insights into search operations and their associated metadata.

## Table Usage Guide

The `algolia_search_metadata` table provides insights into search operations within Algolia. As a developer or data analyst, explore details about search operations through this table, including search queries, hits, processing time, and associated metadata. Utilize it to uncover information about search patterns, performance metrics, and the optimization of search operations.

## Examples

### Get facet data for each index
Analyze the different facets of your search indices to understand their structure and characteristics. This can help in optimizing your search functionality for better user experience.

```sql+postgres
select
  index,
  jsonb_pretty(facets)
from
  algolia_search_metadata;
```

```sql+sqlite
select
  index,
  facets
from
  algolia_search_metadata;
```

### Number of hits and time take per index for a query
Analyze the performance of a specific search term in an Algolia-powered search system. This allows you to understand which indices are most frequently hit and how long it takes to process each hit, helping to optimize search performance.

```sql+postgres
select
  index,
  num_hits,
  processing_time_ms
from
  algolia_search_metadata
where
  query = 'aws vpc'
order by
  processing_time_ms desc;
```

```sql+sqlite
select
  index,
  num_hits,
  processing_time_ms
from
  algolia_search_metadata
where
  query = 'aws vpc'
order by
  processing_time_ms desc;
```