build:
	docker-compose build users-app

run:
	docker-compose up users-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' up