# Story Protocol Backend  

Story Protocol backend provides the protocol data indexing and API services for other products and services 

## Onboarding

### Tool Installation
* Install Make: Run `xcode-select --install`
* Install Docker: Follow instructions at [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
* Install Golang: Guide available [here](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)
* Install Proto: Execute `go get -u github.com/golang/protobuf/{proto,proto-gen-go}`
* Install Protoc-Gen-Go (Use `go install` when outside of a module): `go install github.com/golang/protobuf/protoc-gen-go@latest`
* Install AWS CLI: Instructions at [AWS CLI Installation](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* Install AWS IAM Authenticator: Refer to [AWS IAM Authenticator Installation](https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html)
* Install Kubectl: Guide available at [Kubernetes Tools](https://kubernetes.io/docs/tasks/tools/)
* Install Kubectx: Run `brew install kubectx`
* Install K9s: Execute `brew install k9s`
* Install Typescript: `npm install typescript -g`

### Infrastructure Setup
* Request an admin to: 
  1. Create a new profile for you.
  2. Grant access to the Kubernetes (k8s) cluster.
* Set up your profile using `aws configure set`.
* Log in to SSO using `aws sso login`.
* Verify ECR access with `make ecr-auth`. You should see "Login Succeeded".
* Set up EKS locally (details to be provided).
* Verify EKS access with `k9s`. You should see cluster information without errors.

### Web3-Gateway
The Web3-Gateway service is located in `project-nova/web3-gateway`.
* Install PNPM: `npm install -g pnpm`
* Install TypeScript: `pnpm add typescript -D`
* Run `pnpm install` in the root directory of Web3-Gateway.
* Create a `local.yaml` file in `web3-gateway/config` with the following content:
    ```
    env: local
    wallet_key: <request a wallet key>
    server:
      port: 10002
    ```
* Return to the root directory of `project-nova` and run `make build-proto` to generate proto files.
* Start the server with `make runweb3gateway`.

* Steps to release a new build for Web3-Gateway service:
  1. Build a Docker image with `make build-web3-gateway` (run `make ecr-auth` first for access).
  2. Push the Docker image to ECR using `make push-web3-gateway`.
  3. Deploy the Docker image to the EKS cluster with `make deploy-web3-gateway`.

### API Service
The API service is located in `project-nova/api`.
* Create `local.yaml` and `secrets.yaml` in `project-nova/api/config`. Obtain `local.yaml` contents from a teammate and leave `secrets.yaml` empty.
* Set up a local PostgreSQL for the API service: `docker-compose up -d api-database`.
* Set up the database schema: `make db_up`.
* Start the server with `make runserver`. You should see `[GIN-debug] Listening and serving HTTP on :10001`.

* Steps to release a new build:
  1. Commit your local changes to Git.
  2. Run `make build-api`.
  3. Execute `make push-api`.
  4. Update the `newTag` field in [this file](https://github.com/storyprotocol/project-nova-cd/blob/main/deploy/envs/stag/kustomization.yml) on GitHub (edit directly on GitHub, but first obtain repository access).
  5. Argo CD will automatically deploy the new tag after a few minutes.


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

### V1 tables (For prototype)

***nft_allowlist***: Store the allowlist information for a nft collection

***wallet_merkle_proof***: Store the whitelist proof for a specific allowlist 

***story_franchise***: Store the story franchise information

***story_info***: Store the story information

***story_chapter***: Store the story chapter information

***franchise_collection***: Store the relationship between the franchise and the nft collection

***nft_collection***: Index the information of the nft collection

***nft_token***: Index the information of a specific nft

### Demo tables 

***story_content***: Store the story content for a specific story   

### V2 tables (Connected to the protocol)

***story_info_v2***: Index the story information for a specific story in the protocol

***character_info***: Index the character information for a specific character in the protocol   

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change. Details see: [CONTRIBUTING](/CONTRIBUTING.md)

Please make sure to update tests as appropriate.

## License

[MIT License](/LICENSE.md)