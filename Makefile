PWD := $(shell pwd)

API_NAME := ponies

GO_SWAGGER_CONFIG := $(PWD)/server.yaml
GO_SWAGGER_SCHEME := http

lint:
	@golangci-lint run ./... --issues-exit-code 0 --deadline 5m0s

test:
	@go test `go list ./...`

generate: generate_docs generate_server

generate_server:
	@echo "Generating server from swagger"
	@swagger generate server \
		--target=$(PWD)/swagger_gen \
		--spec=$(PWD)/docs_gen/swagger.yaml \
		--server-package=server \
        --api-package=restapi \
        --principal=models.Principal \
        --name=$(API_NAME) \
        --default-scheme=$(GO_SWAGGER_SCHEME) \
        --config-file=$(GO_SWAGGER_CONFIG)
	@go mod tidy

generate_docs: merge_swagger validate_swagger
	@echo "Generating markdown"
	@swagger generate markdown -f $(PWD)/docs_gen/swagger.yaml --output $(PWD)/docs_gen/api.md

validate_swagger:
	@echo "Running $@"
	@swagger validate $(PWD)/docs_gen/swagger.yaml

merge_swagger:
	@echo "Installing swagger-merger" && yarn global add swagger-merger
	@echo "Merging swagger specification files"
	@mkdir -p $(PWD)/docs_gen
	@swagger-merger -i $(PWD)/swagger/index.yaml -o $(PWD)/docs_gen/swagger.yaml

serve:
	@swagger serve -p 8081 ./docs_gen/swagger.yaml

build_docker:
	@docker build -t org-esv:latest -f ./Dockerfile $(SAAS_PATH)

build_go:
	@go build -v -i -o org-esv ./swagger_gen/cmd

build_env_vars:
	$(eval TENANT_NAME := $(shell kubectl get configmap/platform-config -n fr-platform -o json | jq -r .data.TENANT_NAME))
	$(eval AM_URL := "https://openam-${TENANT_NAME}.forgeblocks.com/am")
	$(eval GOOGLE_PROJECT_ID := $(shell gcloud config list --format 'value(core.project)' 2>/dev/null))
	$(eval OAUTH_CLIENT_ID := openidm-resource-server)
	$(eval OAUTH_CLIENT_SECRET := $(shell kubectl get secret rsfilter-resource-server --namespace fr-platform -o json | jq .data.secret -r | base64 -d; echo))

run_docker: build_env_vars
	@docker run -u root -p 8080:8080 -e PORT=8080 -e HOST=0.0.0.0 \
	  -e AM_URL="${AM_URL}" \
	  -e GOOGLE_PROJECT_ID="${GOOGLE_PROJECT_ID}" \
	  -e OAUTH_CLIENT_ID="${OAUTH_CLIENT_ID}" \
	  -e OAUTH_CLIENT_SECRET="${OAUTH_CLIENT_SECRET}" \
		org-esv:latest

run_local: build_env_vars
	@export AM_URL="${AM_URL}" && \
	  export GOOGLE_PROJECT_ID="${GOOGLE_PROJECT_ID}" && \
	  export OAUTH_CLIENT_ID="${OAUTH_CLIENT_ID}" && \
	  export OAUTH_CLIENT_SECRET="${OAUTH_CLIENT_SECRET}" && \
		go run ./swagger_gen/cmd/main.go --scheme=http --port=8080
