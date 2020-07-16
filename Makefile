APPNAME := simple-blockchain-in-go
DEV_IMAGE_NAME := simple-blockchain-dev

.PHONY: build
build:
	@GO111MODULE=on go build -o $(APPNAME)

.PHONY: clean
clean:
	@rm -f $(APPNAME)

# Cleand and check code
.PHONY: fmt
fmt:
	go fmt `go list`/...

.PHONY: lint
lint:
	./ci/golangci-lint run -c ci/config.yml

.PHONY: update-lint
update-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ci/ v1.27.0

# Test code
.PHONY: tests
tests:
	go test -cover ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=c.out -v ./...
	go tool cover -html=c.out -o coverage.html
	open coverage.html

.PHONY: dev-image
dev-image:
	DOCKER_BUILDKIT=1 docker build -f dev.Dockerfile -t $(DEV_IMAGE_NAME) --ssh default .
