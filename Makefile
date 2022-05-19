.PHONY: dev up down seed

dev:
	modd

up:
	migrate -source file://db/migrations -database 'postgres://localhost:5432/lantern-go?sslmode=disable' up

down:
	migrate -source file://db/migrations -database 'postgres://localhost:5432/lantern-go?sslmode=disable' down

logs:
	doctl apps logs $(LANTERN_APP_ID)
	
seed: 
	go test github.com/aflashyrhetoric/lantern-go/seed	
