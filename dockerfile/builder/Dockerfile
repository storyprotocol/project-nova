FROM ubuntu:latest

# apt list -a <package> to see supported versions of a package.
# Install all build dependencies. protobuf-compiler essentially installs the protoc.
RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    unzip \
    curl \
    tar \
    make \
    nodejs \
    npm

RUN curl -OL https://dl.google.com/go/go1.19.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin:/root/go/bin"

# Install protoc specific version. This installs common protos, e.g. google/protobuf/timestamp.proto,
# under /usr/local/include/google and make it available in PATH - later on "protoc -I." will be able
# to find these protos and compile successfully. 
ARG PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# Golang related.
# https://maelvls.dev/go111module-everywhere/
ENV GO111MODULE=on
RUN go install github.com/golang/protobuf/protoc-gen-go@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

RUN npm install -g typescript@latest