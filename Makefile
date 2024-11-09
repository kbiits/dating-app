include .env

run-http:
	go run ./cmd/http/

run-cron:
	go run ./cmd/cron/

migrate-up:
	migrate -path ./migrations \
		-database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
		up

migrate-down:
	migrate -path ./migrations \
		-database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
		down