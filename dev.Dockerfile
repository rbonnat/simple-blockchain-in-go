# Build stage
FROM golang:1.13.4-alpine3.10 as simple-blockchain-build
RUN apk update && apk add --no-cache \
    git \
    bash \
    make

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make build

# Final stage
FROM alpine:3.10 as simple-blockchain-dev
RUN apk update && apk add --no-cache \
    bash
EXPOSE 8080
WORKDIR /
COPY --from=simple-blockchain-build /app /
CMD ["/simple-blockchain-in-go"]