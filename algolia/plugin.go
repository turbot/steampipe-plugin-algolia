package algolia

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-algolia",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"algolia_api_key":         tableAlgoliaAPIKey(ctx),
			"algolia_index":           tableAlgoliaIndex(ctx),
			"algolia_log":             tableAlgoliaLog(ctx),
			"algolia_search":          tableAlgoliaSearch(ctx),
			"algolia_search_metadata": tableAlgoliaSearchMetadata(ctx),
		},
	}
	return p
}
