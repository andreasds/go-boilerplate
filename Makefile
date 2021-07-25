.PHONY: docker_dev_start
docker_dev_start:
	docker-compose -f docker-compose.dev.yaml up --build

.PHONY: docker_dev_stop
docker_dev_stop:
	docker-compose -f docker-compose.dev.yaml down
