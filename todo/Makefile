build:
 go build -o music-library ./cmd/app/main.go

run:
 ./music-library

migrate:
 migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
 swag init -g /cmd/app/main.go
