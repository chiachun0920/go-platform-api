run:
	docker-compose up -d mongo
	go run ./cmd/web

build:
	docker-compose build apiserver
