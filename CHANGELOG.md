## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#32](https://github.com/turbot/steampipe-plugin-algolia/pull/32))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#32](https://github.com/turbot/steampipe-plugin-algolia/pull/32))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#30](https://github.com/turbot/steampipe-plugin-algolia/pull/30))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#30](https://github.com/turbot/steampipe-plugin-algolia/pull/30))

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#26](https://github.com/turbot/steampipe-plugin-algolia/pull/26))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#26](https://github.com/turbot/steampipe-plugin-algolia/pull/26))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-algolia/blob/main/docs/LICENSE). ([#26](https://github.com/turbot/steampipe-plugin-algolia/pull/26))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#25](https://github.com/turbot/steampipe-plugin-algolia/pull/25))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#18](https://github.com/turbot/steampipe-plugin-algolia/pull/18))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#16](https://github.com/turbot/steampipe-plugin-algolia/pull/16))
- Recompiled plugin with Go version `1.21`. ([#16](https://github.com/turbot/steampipe-plugin-algolia/pull/16))

## v0.3.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#11](https://github.com/turbot/steampipe-plugin-algolia/pull/11))

## v0.2.0 [2022-09-28]

_Bug fixes_

- Fixed the link to Algolia app/API settings page in `docs/index.md` file. ([#7](https://github.com/turbot/steampipe-plugin-algolia/pull/7))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#8](https://github.com/turbot/steampipe-plugin-algolia/pull/8))
- Recompiled plugin with Go version `1.19`. ([#8](https://github.com/turbot/steampipe-plugin-algolia/pull/8))

## v0.1.0 [2022-04-27]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#3](https://github.com/turbot/steampipe-plugin-algolia/pull/3))
- Added support for native Linux ARM and Mac M1 builds. ([#4](https://github.com/turbot/steampipe-plugin-algolia/pull/4))

## v0.0.1 [2022-02-16]

_What's new?_

- New tables added
  - [algolia_api_key](https://hub.steampipe.io/plugins/turbot/algolia/tables/algolia_api_key)
  - [algolia_index](https://hub.steampipe.io/plugins/turbot/algolia/tables/algolia_index)
  - [algolia_log](https://hub.steampipe.io/plugins/turbot/algolia/tables/algolia_log)
  - [algolia_search](https://hub.steampipe.io/plugins/turbot/algolia/tables/algolia_search)
  - [algolia_search_metadata](https://hub.steampipe.io/plugins/turbot/algolia/tables/algolia_search_metadata)
