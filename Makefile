

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
compose-up:
	docker compose -f scripts/docker-compose.yml up -d


pr:
	echo $2

tag:
	$(eval ver = $(shell git describe --abbrev=0 --tags))
	sed -i 's/server:latest/server:$(ver)/' scripts/docker-compose.yml
	


clean:
	rm -rf build



.PHONY: all test clean build
