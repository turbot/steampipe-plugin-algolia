---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/algolia.svg"
brand_color: "#5468FF"
display_name: Algolia
name: algolia
description: Steampipe plugin for querying Algolia indexes, logs and more.
og_description: Query Algolia with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/algolia-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Algolia + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Algolia](https://algolia.com) provides search as a service, offering web search across a client's website using an externally hosted search engine.

For example:

```sql
select
  rank,
  hit
from
  algolia_search
where
  index = 'offices'
  and query = 'usa'
order by
  rank;
```

```
+------+---------------------------------------------------------+
| rank | hit                                                     |
+------+---------------------------------------------------------+
| 1    | { "name": "Scranton", "state": "PA", "country": "USA" } |
| 2    | { "name": "Stamford", "state": "CT", "country": "USA" } |
+------+---------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/algolia/tables)**

## Get started

### Install

Download and install the latest Algolia plugin:

```bash
steampipe plugin install algolia
```

### Configuration

Installing the latest algolia plugin will create a config file (`~/.steampipe/config/algolia.spc`) with a single connection named `algolia`:

```hcl
connection "algolia" {
  plugin  = "algolia"
  app_id  = "D902ND8AK3"
  api_key = "dba7ac3a72a41306c5b34207b9fd9d95"
}
```

* `app_id` - Unique application ID, available in your [Algolia settings](https://www.algolia.com/account/api-keys).
* `api_key` - API key, available in your [Algolia settings](https://www.algolia.com/account/api-keys).

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-algolia
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
