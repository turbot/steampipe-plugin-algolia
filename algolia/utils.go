package algolia

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*search.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "algolia"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*search.Client), nil
	}

	// Default to using env vars
	appID := os.Getenv("ALGOLIA_APP_ID")
	apiKey := os.Getenv("ALGOLIA_API_KEY")

	// But prefer the config
	algoliaConfig := GetConfig(d.Connection)

	if algoliaConfig.AppID != nil {
		appID = *algoliaConfig.AppID
	}
	if algoliaConfig.APIKey != nil {
		apiKey = *algoliaConfig.APIKey
	}

	if appID == "" || apiKey == "" {
		// Credentials not set
		return nil, errors.New("app_id and api_key must be configured")
	}

	conn := search.NewClient(appID, apiKey)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

// Commenting out the function below since it is not used at present. There is a possibility for future usage.

// func unixTimestampToDateTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := int64(d.Value.(int))
// 	if i == 0 {
// 		return nil, nil
// 	}
// 	ts := time.Unix(i, 0)
// 	return ts, nil
// }

func queryParamsToJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	q := d.Value.(string)
	m, err := url.ParseQuery(q)
	if err != nil {
		return nil, err
	}
	js, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(js), nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Algolia API error [404]")
}
