SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/*.go

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
		swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

Install-swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models	

mock: ##make mock files using mockgen
	mockgen -source=pkg/repository/interfaces/user.go -destination=pkg/mock/mockrepo/user_mock.go -package=mockrepo
	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/mock/mockusecase/user_mock.go -package=mockusecase
	mockgen -source=pkg/repository/interfaces/helper.go -destination=pkg/mock/mockrepo/helper_mock.go -package=mockrepo
	mockgen -source=pkg/repository/interfaces/order.go -destination=pkg/mock/mockrepo/order_mock.go -package=mockrepo
	mockgen -source=pkg/helper/userhelper.go -destination=pkg/mock/mockhelper/helper_pkg_mock.go -package=mockrepo
