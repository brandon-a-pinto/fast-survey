build:
	@go build -o bin/fast-survey -C cmd/api
	@chmod +x cmd/api/bin/fast-survey

prepare:
	@cp -r ./examples/.env .
	@go mod tidy

up:
	docker-compose up -d

down:
	docker-compose down
