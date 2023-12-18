FROM golang:1.19.4 as build

WORKDIR /build/

COPY go.mod go.sum ./
RUN go mod download

COPY pkg pkg
COPY api api
COPY Makefile ./

RUN make buildserver

EXPOSE 8090

FROM alpine

WORKDIR /app
COPY --from=build /build/api /app/api
CMD ["/app/api/build/server", "--config=/app/api/config/base.yaml,/app/api/config/prod.yaml"]