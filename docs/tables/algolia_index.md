# Table: algolia_index

Run a index query, each row is a result.

Note: A `query` must be provided in all queries to this table.

## Examples

### List all indices

```sql
select
  name,
  created_at,
  entries
from
  algolia_index
order by
  name
```

### Top 5 largest indices

```sql
select
  name,
  entries,
  data_size
from
  algolia_index
order by
  entries desc
```

### Indices last updated more than 7 days ago

```sql
select
  name,
  updated_at
from
  algolia_index
where
  age(updated_at) > interval '7 days'
```
