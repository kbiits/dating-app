include .env
export

run-http:
	go run ./cmd/http/

run-cron:
	go run ./cmd/cron/

migrate-down-and-up: migrate-down migrate-up

migrate-up:
	migrate -path ./migrations \
		-database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
		up

migrate-down:
	migrate -path ./migrations \
		-database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
		down

generate-integration-test-data:
	node ./integration_test_data_generator.js > ./integration-test-data.json