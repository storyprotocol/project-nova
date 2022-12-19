.PHONY: server db_new

buildserver:
	cd api-server && go build -o build/server cmd/main.go
runserver:
	cd api-server && ./build/server
server:
	make buildserver && make runserver
db_new: 
	migrate create -ext sql -dir api-server/migrations -seq "migration_step_please_change_name"