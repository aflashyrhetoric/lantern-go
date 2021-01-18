package main

import (
	"log"

	"github.com/aflashyrhetoric/lantern-api/db"
	"github.com/aflashyrhetoric/lantern-api/handlers/person"
	"github.com/gin-gonic/gin"

	_ "net/http"
)

func main() {
	r := gin.Default()

	// Use the InitDB function to initialise the global variable.
	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/people", person.GetPeople)
	r.Run() // listen and serve on localhost:8080
}
