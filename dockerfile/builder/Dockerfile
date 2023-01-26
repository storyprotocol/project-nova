FROM ubuntu:latest

# apt list -a <package> to see supported versions of a package.
# Install all build dependencies. protobuf-compiler essentially installs the protoc.
RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    unzip \
    curl \
    tar \
    make 

RUN curl -OL https://dl.google.com/go/go1.19.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin:/root/go/bin"


# Golang related.
# https://maelvls.dev/go111module-everywhere/
ENV GO111MODULE=on