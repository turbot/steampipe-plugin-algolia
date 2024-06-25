package algolia

import (
	"context"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlgoliaSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "algolia_search",
		Description: "Search results (hits) for a given query.",
		List: &plugin.ListConfig{
			ParentHydrate: listIndex,
			Hydrate:       listQueryResult,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "index", Require: plugin.Optional},
				{Name: "query", Require: plugin.Optional},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "index", Type: proto.ColumnType_STRING, Description: "Name of the index for the search result."},
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank (position) of the search result. The top result is number 1."},
			{Name: "hit", Type: proto.ColumnType_JSON, Description: "Hit data of the search result."},
			// Other columns
			{Name: "highlight_result", Type: proto.ColumnType_JSON, Description: "Highlight information."},
			{Name: "object_id", Type: proto.ColumnType_STRING, Description: "Object ID for this search result."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query used to obtain the search result."},
			{Name: "ranking_info", Type: proto.ColumnType_JSON, Description: "Ranking information for the search result."},
			{Name: "snippet_result", Type: proto.ColumnType_JSON, Description: "Snippet information."},
		}),
	}
}

type hitRow struct {
	Index           string      `json:"index"`
	Rank            int         `json:"rank"`
	ObjectID        string      `json:"objectID"`
	Hit             interface{} `json:"hit"`
	RankingInfo     interface{} `json:"_rankingInfo"`
	SnippetResult   interface{} `json:"_snippetResult"`
	HighlightResult interface{} `json:"_highlightResult"`
}

func listQueryResult(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
	for i, hit := range result.Hits {
		hitData := map[string]interface{}{}
		for k, v := range hit {
			switch k {
			// _distinctSeqID is not extracted since it's only present when using
			// the Distinct parameter, which we do not use by default.
			// Common fields defined at https://www.algolia.com/doc/api-reference/api-methods/search/#method-response-hits
			case "objectID", "_rankingInfo", "_snippetResult", "_highlightResult":
				continue
			default:
				hitData[k] = v
			}
		}
		row := hitRow{
			Index:           parentIndex.Name,
			Rank:            i + 1,
			Hit:             hitData,
			ObjectID:        hit["objectID"].(string),
			RankingInfo:     hit["_rankingInfo"],
			SnippetResult:   hit["_snippetResult"],
			HighlightResult: hit["_highlightResult"],
		}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
