package algolia

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "app_id",
			Description: "Unique identifier for the algolia application.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getApplicationId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getApplicationIdMemoized = plugin.HydrateFunc(getApplicationIdUncached).Memoize(memoize.WithCacheKeyFunction(getApplicationIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getApplicationId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getApplicationIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getApplicationIdCacheKey.
func getApplicationIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getApplicationId"
	return key, nil
}

func getApplicationIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cfg := GetConfig(d.Connection)

	return cfg.AppID, nil
}
