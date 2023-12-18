OUTPUT ?= $(PWD)/_output

LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ECR ?= 243963068353.dkr.ecr.us-east-2.amazonaws.com
API_SERVER_IMAGE ?= api
STREAMER_IMAGE ?= streamer

.PHONY: deploy
deploy:
	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${API_SERVER_IMAGE}=${ECR}/${API_SERVER_IMAGE}:${TAG}
	mkdir -p ${OUTPUT}
	kustomize build deploy/overlays/${ENV} -o ${OUTPUT}/deploy.yaml

	cd deploy/overlays/${ENV}; kustomize edit set image ${ECR}/${API_SERVER_IMAGE}=${ECR}/${API_SERVER_IMAGE}:latest
	kubectl apply -f ${OUTPUT}/deploy.yaml