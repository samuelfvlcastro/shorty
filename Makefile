start:
	air

start-ui:
	cd ui && npm run dev


start-database:
	docker-compose up -d

migrate-up:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose down

migrate-fix:
	migrate -path database/migrations -database "postgresql://root:toor@127.0.0.1:5432/shorty?sslmode=disable" -verbose force 1
