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

## Data Model

The definition of the postgres data schema is [here](/api/migrations)

V1 tables (For prototype)

***nft_allowlist***: Store the allowlist information for a nft collection

***wallet_merkle_proof***: Store the whitelist proof for a specific allowlist 

***story_franchise***: Store the story franchise information

***story_info***: Store the story information

***story_chapter***: Store the story chapter information

***franchise_collection***: Store the relationship between the franchise and the nft collection

***nft_collection***: Index the information of the nft collection

***nft_token***: Index the information of a specific nft


Demo tables 

***story_content***: Store the story content for a specific story 


V2 tables (Connected to the protocol)

***story_info_v2***: Index the story information for a specific story in the protocol

***character_info***: Index the character information for a specific character in the protocol 

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change. Details see: [CONTRIBUTING](/CONTRIBUTING.md)

Please make sure to update tests as appropriate.

## License

[MIT License](/LICENSE.md)