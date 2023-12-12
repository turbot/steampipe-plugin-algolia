package algolia

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type algoliaConfig struct {
	AppID  *string `hcl:"app_id"`
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &algoliaConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) algoliaConfig {
	if connection == nil || connection.Config == nil {
		return algoliaConfig{}
	}
	config, _ := connection.Config.(algoliaConfig)
	return config
}
