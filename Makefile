all: install

run:
	docker-compose up --build

install:
	go install -v ./cmd/syslog-send

fmt:
	go fmt ./...

test:
	go test ./... -v
