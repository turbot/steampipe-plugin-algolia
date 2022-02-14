package main

import (
	"github.com/turbot/steampipe-plugin-algolia/algolia"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: algolia.Plugin})
}
