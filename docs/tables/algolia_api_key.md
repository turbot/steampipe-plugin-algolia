# Table: algolia_api_key

API Keys defined in the Algolia account.

## Examples

### List all API Keys

```sql
select
  *
from
  algolia_api_key
```

### Access keys more than 90 days old

```sql
select
  key,
  created_at,
  description
from
  algolia_api_key
where
  age(created_at) > interval '90 days'
```

### Access keys with permission to add objects

```sql
select
  key,
  acl,
  description
from
  algolia_api_key
where
  acl ? 'addObject'
```
