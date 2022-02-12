![image](https://hub.steampipe.io/images/plugins/turbot/algolia-social-graphic.png)

# Algolia Plugin for Steampipe

Use SQL to query indexes and metadata from Algolia.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/algolia)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/algolia/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-algolia/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install algolia
```

Run a query:
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
  rank
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-algolia.git
cd steampipe-plugin-algolia
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/algolia.spc
```

Try it!
```
steampipe query
> .inspect algolia
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-prometheus/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Algolia Plugin](https://github.com/turbot/steampipe-plugin-algolia/labels/help%20wanted)
