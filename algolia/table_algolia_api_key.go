package algolia

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlgoliaAPIKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "algolia_api_key",
		Description: "API keys for the algolia account.",
		List: &plugin.ListConfig{
			Hydrate: listAPIKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key"),
			Hydrate:    getAPIKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Value"), Description: "API key value."},
			// Other columns
			{Name: "acl", Type: proto.ColumnType_JSON, Description: "List of permissions the key contains."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date at which the key has been created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the key."},
			{Name: "indexes", Type: proto.ColumnType_JSON, Description: "The list of targeted indices, if any."},
			// Zero is default (unlimited) for max_*, so do not null it out
			{Name: "max_hits_per_query", Type: proto.ColumnType_INT, Transform: transform.FromGo(), Description: "Maximum number of hits this API key can retrieve in one call."},
			{Name: "max_queries_per_ip_per_hour", Type: proto.ColumnType_INT, Transform: transform.FromGo(), Description: "Maximum number of API calls allowed from an IP address per hour."},
			{Name: "query_parameters", Type: proto.ColumnType_STRING, Description: "Parameters added to all searches with this key."},
			{Name: "referers", Type: proto.ColumnType_JSON},
			{Name: "validity", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the date at which the key expires. (0 means it will not expire automatically)."},
		},
	}
}

func listAPIKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	result, err := conn.ListAPIKeys()
	if err != nil {
		return nil, err
	}
	for _, i := range result.Keys {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getAPIKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := d.EqualsQuals["key"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	result, err := conn.GetAPIKey(key)
	return result, err
}
