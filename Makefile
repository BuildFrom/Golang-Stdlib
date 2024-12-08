run:
	go run ./cmd/api/main.go

sql:
	docker run -d --name postgres -td -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:16

accessdb:
	docker exec -it postgres psql -U postgres

verify-checksums:
	go mod verify

test:
	go test -v ./...


# go test -v ./... | grep -v -e "?" -e "container" -e "Container" -e "RUN"

# go test -v -race -cover ./...

# CREATE DATABASE your_database_name;
# \c your_database_name
# \dt to list tables
# \d your_table_name to describe a table
# \q to quit