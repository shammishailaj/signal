.PHONY: assets build dev docker static

DOCKER_IMAGE = "astrocorp/signal_api"

build:
	go build -o dist/signal

static:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/signal

assets:
	cp ../web/analytics.min.js .

dev:
	go run main.go

docker:
	docker build -t $(DOCKER_IMAGE) .
