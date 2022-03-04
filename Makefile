

build: clean
	go build -o build/gserver ./main.go 
run:
	./build/gserver

	

container:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	@echo  "build server:$(ver)"
	docker build -f scripts/Dockerfile  -t server:$(ver) -t  registry.cn-shanghai.aliyuncs.com/fundshow/goblin:$(ver) .

push-ali:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	docker push registry.cn-shanghai.aliyuncs.com/fundshow/goblin:$(ver)

compose:
	docker compose up 

compose-down:
	docker compose -f scripts/docker-compose.yml down
compose-up: update-tag
	docker compose -f scripts/docker-compose.yml up -d
	sed -i 's/server:v[^"]*/server:server-version/' scripts/docker-compose.yml


update-tag:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	sed -i 's/server:server-version/server:$(ver)/' scripts/docker-compose.yml

update-server: update-tag
	docker compose -f scripts/docker-compose.yml stop server
	docker compose -f scripts/docker-compose.yml up -d --build


clean:
	rm -rf build



.PHONY: all test clean build
