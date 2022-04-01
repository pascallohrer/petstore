image_name := svcpetstore

.PHONY: build run test

build:
	@docker build -t ${image_name}:latest -f Dockerfile .

run:
	@docker run --publish 8080:8080 ${image_name}:latest 

test:
	@docker build -t ${image_name}-test:latest -f Dockerfile.test .
	@docker run -v ${PWD}:/go/app ${image_name}-test:latest
