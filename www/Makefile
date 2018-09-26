.PHONY: dev build docker clean install

DOCKER_IMAGE = "astrocorp/signal_www"

build:
	npm run build

install:
	npm install

dev:
	npm run serve

clean:
	rm -rf dist node_modules

docker:
	docker build -t $(DOCKER_IMAGE) .
