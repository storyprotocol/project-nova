OUTPUT ?= $(PWD)/_output

LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ENV ?= staging
REGION ?= us-east-2
ifeq ($(ENV), prod)
	REGION = us-east-1
endif

ECR ?= 243963068353.dkr.ecr.${REGION}.amazonaws.com
WEB3_GATEWAY_IMAGE ?= web3-gateway

PROTO_DESCRIPTOR_FILE_NAME = proto_descriptor.bin

all: build

build-push: build push

.PHONY: lint
lint:
	npm run lint

.PHONY: build-proto
build-proto:
	rm -rf src/proto/*
	protoc --plugin=node_modules/ts-proto/protoc-gen-ts_proto \
		../proto/v1/web3_gateway/*.proto \
		-I.. \
		--ts_proto_out=src/proto \
		--ts_proto_opt=outputServices=grpc-js,env=node,useOptionals=messages,exportCommonSymbols=false,esModuleInterop=true \
		--descriptor_set_out=src/proto/${PROTO_DESCRIPTOR_FILE_NAME} --include_imports

.PHONY: web3-gateway-server
web3-gateway-server: build-proto
	rm -rf dist; npx tsc -b
	cp src/proto/${PROTO_DESCRIPTOR_FILE_NAME} dist/proto
	cp config/* dist/config/

.PHONY: run-server-local
run-server-local: web3-gateway-server
	SERVER_ENV=local node dist/app/main

.PHONY: deploy
deploy:
	@echo "deploying to ${ENV}"
	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${WEB3_GATEWAY_IMAGE}=${ECR}/${WEB3_GATEWAY_IMAGE}:${TAG}
	mkdir -p ${OUTPUT}
	kustomize build deploy/overlays/${ENV} -o ${OUTPUT}/deploy.yaml

	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${WEB3_GATEWAY_IMAGE}=${ECR}/${WEB3_GATEWAY_IMAGE}:latest
	kubectl apply -f ${OUTPUT}/deploy.yaml
