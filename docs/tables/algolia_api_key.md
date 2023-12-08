---
title: "Steampipe Table: algolia_api_key - Query Algolia API Keys using SQL"
description: "Allows users to query Algolia API Keys, specifically the key details and associated permissions, providing insights into key usage patterns and potential security risks."
---

# Table: algolia_api_key - Query Algolia API Keys using SQL

Algolia is a powerful search-as-a-service solution that allows you to build fast and relevant search into your applications. It provides you with the ability to index and search large amounts of data, and to deliver search results to your users in an instant. Algolia API Keys are used to authenticate requests to Algolia's API and manage permissions for different types of operations.

## Table Usage Guide

The `algolia_api_key` table provides insights into API keys within Algolia's search service. As a DevOps engineer, explore key-specific details through this table, including permissions, validity, and associated metadata. Utilize it to uncover information about keys, such as those with unrestricted permissions, the validity period of keys, and the verification of permission settings.

## Examples

### List all API Keys
Explore the list of all API keys to manage access control, monitor usage, and enhance security measures. This allows you to identify potential vulnerabilities or misuse, ensuring the integrity and confidentiality of your data.

```sql+postgres
select
  *
from
  algolia_api_key;
```

```sql+sqlite
select
  *
from
  algolia_api_key;
```

### Access keys more than 90 days old
Discover the segments that have access keys older than 90 days. This is useful to identify potential security risks and ensure timely key rotation.

```sql+postgres
select
  key,
  created_at,
  description
from
  algolia_api_key
where
  age(created_at) > interval '90 days';
```

```sql+sqlite
select
  key,
  created_at,
  description
from
  algolia_api_key
where
  julianday('now') - julianday(created_at) > 90;
```

### Access keys with permission to add objects
Analyze the settings to understand which access keys have the capability to add new objects. This is beneficial in managing access control and ensuring only authorized keys can modify your data.

```sql+postgres
select
  key,
  acl,
  description
from
  algolia_api_key
where
  acl ? 'addObject'
```

```sql+sqlite
Error: SQLite does not support the '?' operator for JSON objects.
```