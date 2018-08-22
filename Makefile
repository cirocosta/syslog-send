all: install

run:
	docker-compose up --build

install:
	go install -v ./cmd/sample-syslog

fmt:
	go fmt ./...

test:
	go test ./... -v
