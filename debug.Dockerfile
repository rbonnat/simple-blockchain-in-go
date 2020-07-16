# Build stage
FROM golang:1.13.4-alpine3.10 as build-debug-env
RUN apk update && apk add --no-cache \
    git \
    make \
    && go get github.com/derekparker/delve/cmd/dlv

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make build

# Final stage
FROM alpine:3.10
EXPOSE 8080 40000

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat \
    bash

WORKDIR /
COPY --from=build-debug-env /app /
COPY --from=build-debug-env /go/bin/dlv /
# Run delve
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "./simple-blockchain-in-go"]