all: build

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: vet fmt
	go build -o bin/logbook logbook.go

.PHONY: run
run: vet fmt
	go run main.go

