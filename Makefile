.SILENT:

go-run:
	go run cmd/avito-test-task/main.go

run:
	docker compose up --build app