FROM 243963068353.dkr.ecr.us-west-2.amazonaws.com/builder:latest as build

WORKDIR /build/

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install --no-install-recommends -y \
  postgresql-client

# Install golang-migrate for db migration
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/local/bin/migrate

COPY api/migrations api/migrations
COPY bastion bastion

CMD ["sh", "-c", "tail -f /dev/null"]