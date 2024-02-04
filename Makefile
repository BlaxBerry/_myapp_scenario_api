APP_NAME = myapp_scenario_api
APP_CMD_DIR = ./cmd
APP_MAIN_PKG_DIR := $(APP_CMD_DIR)/$(APP_NAME)
APP_BINARY_FILE := ${APP_CMD_DIR}/app
# GOPATH := $(shell go env GOPATH)


dev:
	go run $(APP_MAIN_PKG_DIR)/main.go

build:
	go build -o $(APP_BINARY_FILE) $(APP_MAIN_PKG_DIR)/main.go

preview:
	make build && $(APP_BINARY_FILE)

test:
	go test ./internal/...

format:
	for a in `find ./ -name "*.go"`; do \
        go fmt $${a}; \
    done
