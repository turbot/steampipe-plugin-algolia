package algolia

import (
	"context"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlgoliaSearchMetadata(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "algolia_search_metadata",
		Description: "Get metadata for an algolia search, including the number of hits available, facets and other information.",
		List: &plugin.ListConfig{
			ParentHydrate: listIndex,
			Hydrate:       listQueryResultMetadata,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "index", Require: plugin.Optional},
				{Name: "query", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "index", Type: proto.ColumnType_STRING, Description: "Name of the index for the search result."},
			{Name: "num_hits", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.NbHits"), Description: "The number of hits matched by the query."},
			{Name: "page", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.Page"), Description: "Index of the current page (zero-based). See the page search parameter."},
			// Other columns
			{Name: "ab_test_variant_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.ABTestVariantID"), Description: "If a search encounters an index that is being A/B tested, this reports the variant ID of the index used."},
			// Always null? - {Name: "applied_rules", Type: proto.ColumnType_JSON},
			{Name: "around_lat_long", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.AroundLatLng"), Description: "The computed geo location. Format: ${lat},${lng}, where the latitude and longitude are expressed as decimal floating point numbers."},
			{Name: "automatic_radius", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.AutomaticRadius"), Description: "The automatically computed radius."},
			{Name: "exhaustive_facets_count", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Result.ExhaustiveFacetsCount"), Description: "Whether the facet count is exhaustive (true) or approximate (false)."},
			{Name: "exhaustive_num_hits", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Result.ExhaustiveNbHits"), Description: "Whether the nbHits is exhaustive (true) or approximate (false)."},
			// Always null? - {Name: "explain", Type: proto.ColumnType_JSON},
			{Name: "facets", Type: proto.ColumnType_JSON, Transform: transform.FromField("Result.Facets"), Description: "A mapping of each facet name to the corresponding facet counts."},
			{Name: "facets_stats", Type: proto.ColumnType_JSON, Transform: transform.FromField("Result.FacetsStats"), Description: "Statistics for numerical facets."},
			{Name: "hits_per_page", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.HitsPerPage"), Description: "The maximum number of hits returned per page."},
			// Always null? - {Name: "index", Type: proto.ColumnType_STRING},
			{Name: "index_used", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.IndexUsed"), Description: "Index name used for the query. In the case of an A/B test, the targeted index isn’t always the index used by the query."},
			// Deprecated pagination - {Name: "length", Type: proto.ColumnType_INT, Transform: transform.FromField("Length")},
			{Name: "message", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.Message"), Description: "Used to return warnings about the query."},
			{Name: "num_pages", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.NbPages"), Description: "The number of returned pages. Calculation is based on the total number of hits (nbHits) divided by the number of hits per page (hitsPerPage), rounded up to the nearest integer."},
			// Deprecated pagination - {Name: "offset", Type: proto.ColumnType_INT, Transform: transform.FromField("Offset")},
			{Name: "params", Type: proto.ColumnType_JSON, Transform: transform.FromField("Result.Params").Transform(queryParamsToJSON), Description: "Parameters passed to the search."},
			{Name: "parsed_query", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.ParsedQuery").Transform(queryParamsToJSON), Description: "The query string that will be searched, after normalization. Normalization includes removing stop words (if removeStopWords is enabled), and transforming portions of the query string into phrase queries."},
			{Name: "processing_time_ms", Type: proto.ColumnType_INT, Transform: transform.FromField("Result.ProcessingTimeMS"), Description: "Time the server took to process the request, in milliseconds. This does not include network time."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query for the saerch."},
			{Name: "query_after_removal", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.QueryAfterRemoval"), Description: "A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set."},
			{Name: "query_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.QueryID"), Description: "Unique identifier of the search query, to be sent in Insights methods. This identifier links events back to the search query it represents."},
			{Name: "server_used", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result.ServerUsed"), Description: "Actual host name of the server that processed the request."},
			// Deprecated - {Name: "timeout_counts", Type: proto.ColumnType_BOOL},
			// Deprecated - {Name: "timeout_hits", Type: proto.ColumnType_BOOL},
			{Name: "user_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Result.UserData"), Description: "User data results from the search."},
		},
	}
}

type metadataRow struct {
	Index  string
	Result search.QueryRes
}

func listQueryResultMetadata(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	q := quals["query"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	parentIndex := h.Item.(search.IndexRes)
	index := conn.InitIndex(parentIndex.Name)
	result, err := index.Search(
		q,
		//opt.AroundLatLngViaIP(true),
		opt.AttributesToRetrieve("*"),
		opt.AttributesToSnippet("*:20"),
		opt.ClickAnalytics(true),
		opt.Facets("*"),
		opt.GetRankingInfo(true),
		opt.HitsPerPage(1000),
		opt.Page(0),
		opt.ResponseFields("*"),
		opt.SnippetEllipsisText("…"),
	)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, metadataRow{
		Index:  parentIndex.Name,
		Result: result,
	})
	return nil, nil
}
