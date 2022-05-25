# Docker for Postgres
docker-compose up --build

# Test with cover
go test -v -covermode=atomic -coverprofile=cover.out -cover ./...
go tool cover -func=cover.out