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
