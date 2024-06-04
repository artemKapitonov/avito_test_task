.SILENT:

db-up:
	docker compose up db

lint:
	golangci-lint run

run:
	go build -o ./bin/avito-test-task ./cmd/avito-test-task
	./bin/avito-test-task

run-compose:
	docker compose up --build app
