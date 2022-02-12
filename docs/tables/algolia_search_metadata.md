# Table: algolia_search_metadata

Get metadata for an algolia search, including the number of hits available,
facets and other information.

Note:
* Use [algolia_search](../algolia_search) to get search results.

## Examples

### Get facet data for each index

```sql
select
  index,
  jsonb_pretty(facets)
from
  algolia_search_metadata
```

### Number of hits and time take per index for a query

```sql
select
  index,
  num_hits,
  processing_time_ms
from
  algolia_search_metadata
where
  query = 'aws vpc'
order by
  processing_time_ms desc
```
