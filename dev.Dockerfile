FROM golang:1.13.4-alpine3.10 as BUILDER
RUN apk update && apk add --no-cache \
    git \
    bash \
    make

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build

CMD ["/app/simple-blockchain-in-go"]