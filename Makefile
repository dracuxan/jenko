run: build
	@./bin/app

build:
	@go build -ldflags "-s -w" -o bin/app

map:
	@tree ~/logs/Black_Hat_Go/ > README.md

json:
	@go run main.go | jq
