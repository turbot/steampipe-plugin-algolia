package algolia

import (
	"bufio"
	"context"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableAlgoliaLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "algolia_log",
		Description: "Recent log entries across all indexes and operation types. Results are limited to last 1000 entries and/or a maximum of 7 days age.",
		List: &plugin.ListConfig{
			Hydrate: listLog,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the log entry was created."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Transform: transform.FromField("IP"), Description: "IP address of the request client."},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "HTTP method used for the query, e.g. GET, POST."},
			{Name: "answer_code", Type: proto.ColumnType_INT, Description: "Code of the answer."},
			{Name: "index", Type: proto.ColumnType_STRING, Description: "Index the query was executed against, or null for metadata queries."},
			// Other columns
			{Name: "answer", Type: proto.ColumnType_STRING, Description: "Answer body, truncated after 1000 characters. JSON format, but returned as a string due to the truncation."},
			{Name: "exhaustive", Type: proto.ColumnType_BOOL, Description: "Exhaustive flags used during the query."},
			{Name: "inner_queries", Type: proto.ColumnType_JSON, Description: "Contains an object for each performed query with the indexName, queryID, offset, and userToken."},
			{Name: "number_api_calls", Type: proto.ColumnType_INT, Transform: transform.FromField("NbAPICalls"), Description: "Number of API calls."},
			{Name: "processing_time_ms", Type: proto.ColumnType_INT, Transform: transform.FromField("ProcessingTime").Transform(convertDurationToMs), Description: "Processing time for the request in milliseconds."},
			{Name: "query_body", Type: proto.ColumnType_STRING, Description: "Request body, truncated after 1000 characters."},
			{Name: "query_headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("QueryHeaders").Transform(convertHeaders), Description: "HTTP headers for the query."},
			{Name: "query_number_hits", Type: proto.ColumnType_INT, Transform: transform.FromField("QueryNbHits"), Description: "Number of hits returned for the query."},
			{Name: "sha1", Type: proto.ColumnType_STRING, Transform: transform.FromField("SHA1"), Description: "SHA1 ID of the log entry."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "URL of the query request."},
		},
	}
}

func listLog(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	// 1000 is the maximum allowed, and paging is not supported.
	// The API defaults to the most recent log entries across all indices and types.
	result, err := conn.GetLogs(opt.Length(1000))
	if err != nil {
		return nil, err
	}
	for _, i := range result.Logs {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func convertDurationToMs(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dur := d.Value.(time.Duration)
	return dur.Milliseconds(), nil
}

// Convert the headers string to a map[string]string.
func convertHeaders(_ context.Context, d *transform.TransformData) (interface{}, error) {
	headerString := d.Value.(string)
	// don't forget to make certain the headers end with a second "\r\n"
	reader := bufio.NewReader(strings.NewReader(headerString + "\r\n"))
	tp := textproto.NewReader(reader)
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	// http.Header and textproto.MIMEHeader are both just a map[string][]string
	httpHeader := http.Header(mimeHeader)
	return httpHeader, nil
}
