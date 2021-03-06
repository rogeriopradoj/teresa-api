SWAGGER_SPEC=swagger.yml
TERESA_API_NAME=Teresa
IMAGE_NAME=teresa
IMAGE_VERSION=0.1.1
IMAGE_INSTANCE=default
TERESA_DOCKER_PORT ?= 8080

DOCKER_RUN_CMD=docker run \
	-e TERESAK8S_HOST=$(TERESAK8S_HOST) \
	-e TERESAK8S_USERNAME=$(TERESAK8S_USERNAME) \
	-e TERESAK8S_PASSWORD=$(TERESAK8S_PASSWORD) \
	-e TERESAK8S_INSECURE=$(TERESAK8S_INSECURE) \
	-e TERESABUILDER_STORAGE_AWS_KEY=$(TERESABUILDER_STORAGE_AWS_KEY) \
	-e TERESABUILDER_STORAGE_AWS_SECRET=$(TERESABUILDER_STORAGE_AWS_SECRET) \
	-e TERESABUILDER_STORAGE_AWS_REGION=$(TERESABUILDER_STORAGE_AWS_REGION) \
	-e TERESABUILDER_STORAGE_AWS_BUCKET=$(TERESABUILDER_STORAGE_AWS_BUCKET) \
	-e TERESABUILDER_WAIT_POD_TIMEOUT=$(TERESABUILDER_WAIT_POD_TIMEOUT) \
	-e TERESABUILDER_WAIT_LB_TIMEOUT=$(TERESABUILDER_WAIT_LB_TIMEOUT) \
	-p $(TERESA_DOCKER_PORT):8080

help:
	@echo "Targets are:\n"
	@echo "build"
	@echo " build the teresa API server docker image"
	@echo
	@echo "run"
	@echo " run the teresa API docker image"
	@echo
	@echo "start"
	@echo " run the teresa API docker image as a daemon"
	@echo
	@echo "stop"
	@echo " stop the teresa API docker image"
	@echo
	@echo "validate"
	@echo " validate swagger file $(SWAGGER_SPEC) aginst swagger specification 2.0"
	@echo
	@echo "gen-api-server"
	@echo " generate the API server code as described by $(SWAGGER_SPEC)"
	@echo
	@echo "gen-api-client"
	@echo " generate the API client code as described by $(SWAGGER_SPEC)"
	@echo
	@echo "To run the API container you'll have to set the following env variables:"
	@echo
	@echo "	TERESAK8S_HOST"
	@echo "	TERESAK8S_USERNAME"
	@echo "	TERESAK8S_PASSWORD"
	@echo "	TERESAK8S_INSECURE"
	@echo "	TERESABUILDER_STORAGE_AWS_KEY"
	@echo "	TERESABUILDER_STORAGE_AWS_SECRET"
	@echo "	TERESABUILDER_STORAGE_AWS_REGION"
	@echo "	TERESABUILDER_STORAGE_AWS_BUCKET"
	@echo "	TERESABUILDER_WAIT_POD_TIMEOUT"
	@echo "	TERESABUILDER_WAIT_LB_TIMEOUT"
	@echo "	TERESA_DOCKER_PORT - optional, defaults to 8080"
	@echo

all: help

build:
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) .

run:
	$(DOCKER_RUN_CMD) --rm --name $(IMAGE_NAME)-$(IMAGE_INSTANCE) $(IMAGE_NAME):$(IMAGE_VERSION)

start:
	$(DOCKER_RUN_CMD) -d $(IMAGE_NAME):$(IMAGE_VERSION)

stop:
	docker stop $(IMAGE_NAME)-$(IMAGE_INSTANCE)

shell:
	docker run --rm --name $(IMAGE_NAME)-$(IMAGE_INSTANCE) -i -t $(IMAGE_NAME):$(IMAGE_VERSION) /bin/bash

gen-api-server:
	swagger generate server -A $(TERESA_API_NAME) -f $(SWAGGER_SPEC)

run-api-server:
	go run ./cmd/teresa-server/main.go --port 8080

gen-api-client:
	swagger generate client -A $(TERESA_API_NAME) -f $(SWAGGER_SPEC)

validate:
	swagger validate $(SWAGGER_SPEC)

swagger-docs:
	go run docs/webserver.go swagger.yml
