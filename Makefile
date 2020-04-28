.DEFAULT_GOAL := build

all: build

build-code:
	go get -d .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/bookdata-api .

build:
	docker build --no-cache -f ./Dockerfile -t bookdata-api:latest --target=final .

run:
	docker run --rm --name bookdata-api -p 8080:8080 -v $$HOST_PROJECT_PATH/assets:/data --env BOOKS_DATABASE_FILE=/data/books.csv bookdata-api:latest

run-test:
	docker run --rm --name bookdata-api -p 8080:8080 --env BOOKS_DATABASE_FILE=/books.csv bookdata-api:latest