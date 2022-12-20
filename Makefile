LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ECR ?= 243963068353.dkr.ecr.us-west-2.amazonaws.com

BUILDER_IMAGE ?= builder
ECR_BUILDER_IMAGE ?= ${ECR}/${BUILDER_IMAGE}

API_SERVER_IMAGE ?= api-server

DOCKER_BUILD=docker build --cache-from

ecr-auth:
	aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin ${ECR}

buildserver:
	cd api && CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o build/server cmd/main.go

runserver:
	cd api && ./build/server

.PHONY: server
server:
	make buildserver && make runserver

.PHONY: db_new
db_new: 
	migrate create -ext sql -dir api/migrations -seq "migration_step_please_change_name"

.PHONY: builder
builder: ecr-auth
	${DOCKER_BUILD} ${BUILDER_IMAGE} -t ${BUILDER_IMAGE}:${TAG} -t ${ECR_BUILDER_IMAGE}:${TAG} dockerfile/builder/
	docker push ${ECR_BUILDER_IMAGE}:${TAG}
	${DOCKER_BUILD} ${BUILDER_IMAGE} -t ${BUILDER_IMAGE}:latest -t ${ECR_BUILDER_IMAGE}:latest dockerfile/builder/
	docker push ${ECR_BUILDER_IMAGE}:latest

build-%:
	docker-compose build $*

push-api:
	docker tag ${API_SERVER_IMAGE} ${ECR}/${API_SERVER_IMAGE}:${TAG}
	docker tag ${API_SERVER_IMAGE} ${ECR}/${API_SERVER_IMAGE}:latest
	docker push ${ECR}/${API_SERVER_IMAGE}:${TAG}
	docker push ${ECR}/${API_SERVER_IMAGE}:latest	

deploy-%:
	cd $*; make deploy