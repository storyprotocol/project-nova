.PHONY: server

buildserver:
	cd server && go build -o build/server cmd/main.go
runserver:
	cd server && ./build/server
server:
	make buildserver && make runserver