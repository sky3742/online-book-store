migrate:
	go run cmd/db/migrate/migrate.go

rollback:
	go run cmd/db/rollback/rollback.go

seeds:
	go run cmd/db/seeds/seeds.go

start:
	go run cmd/api/main.go

watch:
	air

mock:
	mockery --all --dir internal/service --output internal/service/mock
	mockery --all --dir internal/repository --output internal/repository/mock
