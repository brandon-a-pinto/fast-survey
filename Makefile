build:
	@go build -o bin/fast-survey -C cmd/api

run: build
	@./cmd/api/bin/fast-survey
