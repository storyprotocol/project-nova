# project-nova

```bash
  ecr-auth:            - Authenticate ECR
  builder:             - Build and push builder image
  buildserver:         - Build api server locally
  runserver:           - Run api server locally
  server:              - Build and then run api server locally
  push-api:            - Push local api server image to ECR

  db_new:              - Create new DB migration script for api server
  db_up:               - Apply new DB migration script to local Postgres DB for testing
  db_down:             - Tear down local Postgres DB tables
  db_shell:            - Open local Postgres DB console

  build-{service}:     - Build specific service
  push-{service}:      - Push the current local image for the service to ECR
  deploy-{service}:    - Deploy the specific service using the latest image in the ECR, need to specific environment with ENV
                         For example: ENV=dev make deploy-bastion
  lint:                - Run linter
```
