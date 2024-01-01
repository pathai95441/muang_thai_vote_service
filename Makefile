GOPRIVATE_REPO := github.com/pathai95441

GO_INSTALLED := $(shell which go)
GO_FILES = $(shell go list ./... | grep -v core/specs | grep -v /internal/domain/db_models_gen | grep -v /src/restapis/api_gen |grep -v internal/integration )

OPENAPI_V3_PATTERN_JSON := "^ *\"openapi\": \"?3\.0"
OPENAPI_V3_PATTERN_YAML := "^openapi: ['\"]?3\.0"

all: gen fmt vet service/test

goprivate:
	@go env GOPRIVATE | grep "$(GOPRIVATE_REPO)" > /dev/null || go env -w GOPRIVATE="$(shell go env GOPRIVATE),$(GOPRIVATE_REPO)"

upgrade-all: tidy
	go get -d -u ./...
	go mod tidy
	go mod vendor

tidy: goprivate
	@go mod tidy
	@go mod vendor

tools-install:
ifndef GO_INSTALLED
	$(error "go is not installed")
endif
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0
	@go install github.com/golang/mock/mockgen@v1.6.0
	@go install github.com/cweill/gotests/gotests@v1.6.0
	@go install github.com/onsi/ginkgo/v2/ginkgo@v2.13.0
	@go install github.com/jandelgado/gcov2lcov@v1.0.5

tools-upgrade:
	go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/cweill/gotests/gotests
	go get -u github.com/onsi/ginkgo/v2/ginkgo
	go mod tidy

generate gen: tools-install goprivate
#cleanup
	@rm -rf ./src/restapis/api_gen/*
	@rm -rf ./pkg/vote_service/*
#create folder to store server
	@mkdir -p ./src/restapis/api_gen
	@mkdir -p ./pkg/vote_service
#generate server
	$(shell go env GOPATH)/bin/oapi-codegen --old-config-style -generate types -o ./src/restapis/api_gen/openapi_types.gen.go -package api_gen docs/vote_service.yml
	$(shell go env GOPATH)/bin/oapi-codegen --old-config-style -generate server -o ./src/restapis/api_gen/openapi_api.gen.go -package api_gen docs/vote_service.yml
#generate pkg client
	$(shell go env GOPATH)/bin/oapi-codegen --old-config-style -generate types -o ./pkg/vote_service/vote_service_api_client_types.gen.go -package vote_service docs/vote_service.yml
	$(shell go env GOPATH)/bin/oapi-codegen --old-config-style -generate client -o ./pkg/vote_service/vote_service_api_client.gen.go -package vote_service docs/vote_service.yml
#generate pkg client mock with mockgen
	$(shell go env GOPATH)/bin/mockgen -source=./pkg/vote_service/vote_service_api_client.gen.go -destination=./pkg/vote_service/mock/api_client.gen.go
#generate other mock
	@go generate ./...
#vendor and tidy if any requires
	@echo "Generated server and client..."
	@echo "If error occurred beyond this message, try fixing compilation errors and run \`make tidy\`."
	@go mod tidy
	@go mod vendor
	@echo "Success!"

gen-db:
#generate database models with sqlboiler
	sqlboiler mysql

run: goprivate
	@go run cmd/vote_service/main.go --print-console-format --print-stacktrace --debug-mode --docs-enabled --pubsub-enabled

build: generate
	@rm -rf bin
	@mkdir -p bin
	@go build -o bin/server core/server/main.go
	@echo "Success! Binaries can be found in 'bin' dir"

vet:
	@go vet $(GO_FILES)

fmt:
	@go fmt $(GO_FILES)

test: goprivate
	@go test $(GO_FILES) -cover --race

tests: goprivate
	@ENV=unittest ginkgo -r -cover --skip-package internal/integration,internal/domain/db_models_gen

unit-test: goprivate
	@ENV=unittest ginkgo -r -cover -coverprofile=unitcoverage.coverprofile --junit-report=report.xml --output-dir=./ --skip-package internal/domain/db_models_gen,internal/integration

generate-unitcoverage:
	@gcov2lcov -infile=unitcoverage.coverprofile -outfile=unitcoverage.lcov

generate-intcoverage:
	@gcov2lcov -infile=integrationcoverage.coverprofile -outfile=integrationcoverage.lcov
	
open-unitcoverage: unit-test generate-unitcoverage
	@genhtml -o coverage unitcoverage.lcov
	@open coverage/index.html

compose-dev:
	@docker-compose \
		-f docker-compose.yml \
		up --build

compose-down:
	@docker-compose \
		-f docker-compose.yml \
		down

dbmate:
	@docker run --rm \
		--platform linux/x86_64 \
		--network=host \
		-e DATABASE_URL=mysql://root:password@127.0.0.1:3306/vote \
		-e DBMATE_SCHEMA_FILE=/data/schema.sql \
		-e DBMATE_MIGRATIONS_DIR=/migrations \
		-v $(PWD)/db:/data \
		-v $(PWD)/db/migrations:/migrations \
		ghcr.io/amacneil/dbmate \
			$(ARGS)
ifndef ARGS
	@echo
	@echo 'Missing argument, try `make dbmate ARGS="foo bar"` for `dbmate foo bar`'
	@echo
endif
