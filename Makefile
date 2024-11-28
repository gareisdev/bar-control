.PHONY: run

build:
	@go build -o ./bin/barcontrol cmd/main/main.go 

run: build
	./bin/barcontrol

run-dev:
	@air .
