DB_URL=postgres://postgres:my_password@localhost:54320/ks_aggregator_rates?sslmode=disable

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

build:
	go build

run:
	go run .

.PHONY: migrateup migratedown run
