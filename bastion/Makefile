OUTPUT ?= $(PWD)/_output

LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ENV ?= staging
REGION ?= us-east-2
ifeq ($(ENV), prod)
	REGION = us-east-1
endif

ECR ?= 243963068353.dkr.ecr.${REGION}.amazonaws.com
BASTION_IMAGE ?= bastion

.PHONY: deploy
deploy:
	@echo "deploying to ${ENV}"
	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${BASTION_IMAGE}=${ECR}/${BASTION_IMAGE}:${TAG}
	mkdir -p ${OUTPUT}
	kustomize build deploy/overlays/${ENV} -o ${OUTPUT}/deploy.yaml

	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${BASTION_IMAGE}=${ECR}/${BASTION_IMAGE}:latest
	kubectl apply -f ${OUTPUT}/deploy.yaml

.PHONY: build-task
build-task:
	CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o /usr/local/bin/task task/main.go