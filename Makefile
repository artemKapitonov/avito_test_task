.SILENT:

lint:
	golangci-lint run

go-build:
	go build -o ./bin/avito-test-task ./cmd/avito-test-task
	./bin/avito-test-task

run-compose:
	docker compose up --build app