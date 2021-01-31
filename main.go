package main

import (
	"log"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "net/http"
)

func main() {
	r := gin.Default()

	// Add some middleware for cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	r.Use(cors.New(config))

	// Use the InitDB function to initialise the global variable.
	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Routes
	r.GET("/people", handlers.GetPeople)
	r.POST("/people", handlers.CreatePerson)
	r.GET("/people/:id", handlers.ShowPerson)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)

	// Notes
	r.POST("/notes/:person_id", handlers.CreateNote)
	r.DELETE("/notes/:id", handlers.DeleteNote)

	// Pressure Points
	r.POST("/pressure-points/:person_id", handlers.CreatePressurePoint)
	r.DELETE("/pressure-points/:id", handlers.DeletePressurePoint)

	// Run
	r.Run() // listen and serve on localhost:8080
}
