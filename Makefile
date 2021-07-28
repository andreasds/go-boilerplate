.PHONY: clean
clean:
	@find . -name *mock* -delete

.PHONY: dev
dev:
	go run ./cmd/go-api-server/api_server.go

.PHONY: docker_dev_start
docker_dev_start:
	docker-compose -f docker-compose.dev.yaml up --build

.PHONY: docker_dev_stop
docker_dev_stop:
	docker-compose -f docker-compose.dev.yaml down

.PHONY: docker_test
docker_test:
	docker-compose -f docker-compose.test.yaml up --build
	docker-compose -f docker-compose.test.yaml down

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test: clean generate
	go test -v ./test/...
