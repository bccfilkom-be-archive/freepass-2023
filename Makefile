postgres:
	podman run -h pg-server --name postgres-server -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine3.17
	
createdb:
	podman exec -ti postgres-server createdb --username=root --owner=root bcc_university

dropdb:
	podman exec -ti postgres-server dropdb bcc_university

server:
	go run main.go

.PHONY: postgres createdb dropdb server