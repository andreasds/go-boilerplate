.PHONY: clean
clean:
	@find . -name *mock* -delete
	@rm -rf tools/wire_gen.go

.PHONY: dev
dev:
	go run github.com/google/wire/cmd/wire tools/wire.go
	go run ./cmd/go-api-server/api_server.go

.PHONY: docker_dev_start
docker_dev_start: generate
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
