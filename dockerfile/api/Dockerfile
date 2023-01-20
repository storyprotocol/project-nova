ARG builder_image_tag
FROM 243963068353.dkr.ecr.us-west-2.amazonaws.com/builder:${builder_image_tag} as build

WORKDIR /build/

COPY go.mod go.sum ./
RUN go mod download

COPY pkg pkg
COPY api api
COPY Makefile ./

RUN make buildserver

EXPOSE 8090

# See https://docs.docker.com/develop/develop-images/dockerfile_best-practices/

FROM alpine

WORKDIR /var/log
RUN mkdir -p /var/log/storyprotocol

WORKDIR /app
COPY --from=build /build/api /app/api
CMD ["/app/api/build/server", "--config=/app/api/config/prod.yaml"]
