start:
	go run cmd/server/main.go
redis:
	docker compose up --build -d