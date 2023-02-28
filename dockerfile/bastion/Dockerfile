ARG DOCKERHUB_URL
FROM ${DOCKERHUB_URL}builder:latest as build

WORKDIR /build/

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install --no-install-recommends -y \
  postgresql-client

# Install golang-migrate for db migration
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/local/bin/migrate

RUN apt-get install --no-install-recommends -y jq

COPY go.mod go.sum ./
RUN go mod download

COPY api/migrations api/migrations
COPY script/dropdb.sh script/dropdb.sh
COPY bastion bastion
COPY pkg pkg

# Build task runner app 
RUN cd bastion && make build-task

CMD ["sh", "-c", "tail -f /dev/null"]