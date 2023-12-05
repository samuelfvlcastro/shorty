start:
	air

start-ui:
	cd ui && npm run dev

migrate_up:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose up

migrate_down:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose down

migrate_fix:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose force 1
