.PHONY: dev up down seed

dev:
	modd

up:
	migrate -source file://db/migrations -database 'postgres://localhost:5432/lantern-go?sslmode=disable' up

down:
	migrate -source file://db/migrations -database 'postgres://localhost:5432/lantern-go?sslmode=disable' down
	
seed: 
	go test github.com/aflashyrhetoric/lantern-go/seed	
