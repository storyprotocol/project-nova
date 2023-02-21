LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ECR ?= 243963068353.dkr.ecr.us-west-2.amazonaws.com

BUILDER_IMAGE ?= builder
ECR_BUILDER_IMAGE ?= ${ECR}/${BUILDER_IMAGE}

API_SERVER_IMAGE ?= api-server

DOCKER_BUILD=docker build --cache-from

DEVELOPMENT_DB_URI = postgresql://postgres:@api-database:5432/postgres?sslmode=disable

help:
	@echo '  ecr-auth:            - Authenticate ECR'
	@echo '  builder:             - Build and push builder image'
	@echo '  buildserver:         - Build api server locally'
	@echo '  runserver:           - Run api server locally'
	@echo '  server:              - Build and then run api server locally'
	@echo '  buildstreamer:       - Build streamer locally'
	@echo '  runstreamer:         - Run streamer locally'
	@echo '  streamer:            - Build and then run streamer locally'
	@echo '  push-api:            - Push local api server image to ECR'
	@echo ''
	@echo '  db_new:              - Create new DB migration script for api server'
	@echo '  db_up:               - Apply new DB migration script to local Postgres DB for testing'
	@echo '  db_down:             - Tear down local Postgres DB tables'
	@echo '  db_shell:            - Open local Postgres DB console'
	@echo ''
	@echo '  build-{service}:     - Build specific service'
	@echo '  push-{service}:      - Push the current local image for the service to ECR'
	@echo '  deploy-{service}:    - Deploy the specific service using the latest image in the ECR, need to specific environment with ENV'
	@echo '                         For example: ENV=dev make deploy-bastion'
	@echo '  restart-api          - Restart API server deployment'
	@echo '  lint:                - Run linter'
	@echo '  abigen:              - Create golang abi client for smart contracts based on the input json file.'
	@echo '                         For example: make abigen package=erc721. package corresponding to the input json name and output package name'

ecr-auth:
	aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin ${ECR}

buildserver:
	cd api && CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o build/server cmd/api/main.go

runserver:
	cd api && ./build/server --config=config/local.yaml,config/secrets.yaml

.PHONY: server
server:
	make buildserver && make runserver

buildstreamer:
	cd api && CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o build/streamer cmd/streamer/main.go

runstreamer:
	cd api && ./build/streamer --config=config/streamer/local.yaml,config/streamer/secrets.yaml

.PHONY: streamer
streamer:
	make buildstreamer && make runstreamer	

preparedb:
	docker compose up -d bastion api-database

.PHONY: db_new
db_new: 
	migrate create -ext sql -dir api/migrations -seq "migration_step_please_change_name"

.PHONY: db_up
db_up: preparedb
	docker exec project-nova-bastion-1 migrate -database ${DEVELOPMENT_DB_URI} -path /build/api/migrations -verbose up

.PHONY: db_down
db_down: preparedb
	docker exec -e DATABASE_URI=${DEVELOPMENT_DB_URI} project-nova-bastion-1 sh /build/script/dropdb.sh 	

PHONY: db_shell
db_shell: preparedb
	docker exec -it project-nova-bastion-1 psql ${DEVELOPMENT_DB_URI} 

.PHONY: builder
builder: ecr-auth
	${DOCKER_BUILD} ${BUILDER_IMAGE} -t ${BUILDER_IMAGE}:${TAG} -t ${ECR_BUILDER_IMAGE}:${TAG} dockerfile/builder/
	docker push ${ECR_BUILDER_IMAGE}:${TAG}
	${DOCKER_BUILD} ${BUILDER_IMAGE} -t ${BUILDER_IMAGE}:latest -t ${ECR_BUILDER_IMAGE}:latest dockerfile/builder/
	docker push ${ECR_BUILDER_IMAGE}:latest

build-%:
	docker-compose build $*

push-%: ecr-auth
	docker tag $* ${ECR}/$*:${TAG}
	docker tag $* ${ECR}/$*:latest
	docker push ${ECR}/$*:${TAG}
	docker push ${ECR}/$*:latest	

push-api: ecr-auth
	docker tag ${API_SERVER_IMAGE} ${ECR}/${API_SERVER_IMAGE}:${TAG}
	docker tag ${API_SERVER_IMAGE} ${ECR}/${API_SERVER_IMAGE}:latest
	docker push ${ECR}/${API_SERVER_IMAGE}:${TAG}
	docker push ${ECR}/${API_SERVER_IMAGE}:latest	

deploy-%:
	cd $*; ENV=${ENV} make deploy

deploy-streamer:
	cd api; ENV=${ENV} make deploy-streamer

restart-api:
	kubectl rollout restart deployment api-server -n edge

lint:
	golangci-lint run

.PHONY: abigen
abigen:
	mkdir -p ./pkg/abi/${package}; abigen --abi=./resource/abi/${package}.json --pkg=${package} --out=./pkg/abi/${package}/${package}.go