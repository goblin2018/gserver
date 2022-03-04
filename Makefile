
dev-init:
	docker compose -f scripts/docker-compose-test.yml up -d


container:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	@echo  "build server:$(ver)"
	docker build -f scripts/Dockerfile  -t server:$(ver) -t  registry.cn-shanghai.aliyuncs.com/fundshow/goblin:$(ver) .

push-ali:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	docker push registry.cn-shanghai.aliyuncs.com/fundshow/goblin:$(ver)

compose-up: update-tag
	docker compose -f scripts/docker-compose.yml up -d
	sed -i 's/server:v[^"]*/server:server-version/' scripts/docker-compose.yml


update-tag:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	sed -i 's/server:server-version/server:$(ver)/' scripts/docker-compose.yml

update-server: update-tag
	docker compose -f scripts/docker-compose.yml stop server
	docker compose -f scripts/docker-compose.yml up -d --build


dev:
  cat $(shell go run main.go -e development)


clean:
	rm -rf build



.PHONY: all test clean build
