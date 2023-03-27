LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

ENV ?= staging
REGION ?= us-east-2
ifeq ($(ENV), prod)
	REGION = us-east-1
endif

ECR ?= 243963068353.dkr.ecr.${REGION}.amazonaws.com

BUILDER_IMAGE ?= builder
ECR_BUILDER_IMAGE ?= ${ECR}/${BUILDER_IMAGE}

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
	@echo '  restart-{service}    - Restart specific service deployment'
	@echo '  lint:                - Run linter'
	@echo '  abigen:              - Create golang abi client for smart contracts based on the input json file.'
	@echo '                         For example: make abigen package=erc721. package corresponding to the input json name and output package name'
	@echo '  s3-download:         - S3 download based on the project and chapter and env passed in.'
	@echo '                         For example: make s3 dowload project=project-nova chapter=1:1:1 env=staging' 
	@echo '  s3-upload:           - S3 upload based on the project and chapter and env passed in.'

ecr-auth:
	aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR}

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
	@echo "pushing to ${ENV}"
	docker tag $* ${ECR}/$*:${TAG}
	docker tag $* ${ECR}/$*:latest
	docker push ${ECR}/$*:${TAG}
	docker push ${ECR}/$*:latest	

deploy-%:
	cd $*; ENV=${ENV} make deploy

deploy-streamer:
	cd api; ENV=${ENV} make deploy-streamer

restart-%:
	kubectl rollout restart deployment $* -n edge

lint:
	golangci-lint run

.PHONY: abigen
abigen:
	mkdir -p ./pkg/abi/${package}; abigen --abi=./resource/abi/${package}.json --pkg=${package} --out=./pkg/abi/${package}/${package}.go

.PHONY: s3-upload
s3-upload:
	aws s3 cp api/resource/content/${project}/${chapter}/content.json  s3://${project}-content-${env}/${chapter}/content.json

.PHONY: s3-download
s3-download:
	aws s3 cp s3://${project}-content-${env}/${chapter}/content.json api/resource/content/${project}/${chapter}/content.json  