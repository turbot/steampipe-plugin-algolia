# Table: algolia_search

Run a search query, each row is a result.

Notes:
* Use [algolia_search_metadata](../algolia_search_metadata) to get search result metadata.

## Examples

### Search the offices index for the term "usa"

```sql
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
  rank
```

### Get all records from all indexes

Warning: This is a very large query, use carefully!

```sql
select
  *
from
  algolia_search
```
