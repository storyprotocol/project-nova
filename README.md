# Story Protocol Backend  

Story Protocol backend provides the protocol data indexing and API services for other products and services 

## Prerequisites
- Install Docker
- Install Go

## Quick Start 
The Makefile handles most of the local development operation tasks

**Build and run API server**

`make server` 

**Build and run streamer**

`make streamer`

**Generate code from proto interface**

`make build-proto`

**Generate code from abi interface**

`make abigen package=<package name>`

## Database Operation

Make sure the docker is installed in local and docker compose up is run before performing db operations.

Database schema definition is stored in api/migrations. 

To add a pair of new migration files, 

- make db_new

To upgrade the db tables based on the migration

- make db_up

To connect to the db shell

- make db_shell

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change. Details see: [CONTRIBUTING](/CONTRIBUTING.md)

Please make sure to update tests as appropriate.

## License

[MIT License](/LICENSE.md)