STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/algolia@latest/steampipe-plugin-algolia.plugin -tags "${BUILD_TAGS}" *.go
