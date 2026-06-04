DB_URL=sqlite3://./data/scheduler.db
MIGRATE=go run github.com/golang-migrate/migrate/v4/cmd/migrate@latest
MIGRATIONS=internal/database/migrations
build:
	go build -o uptique
serve:
	./uptique serve
build_and_serve:
	make build && make serve
migration_create:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS) -seq $(name)

migrate-up:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
        -database "sqlite3://./data/scheduler.db" -path internal/database/migrations up

migrate-down:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
        -database "sqlite3://./data/scheduler.db" -path internal/database/migrations down 1
migrate-force:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
		-database "sqlite3://./data/scheduler.db" -path internal/database/migrations force 1