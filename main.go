package main

import (
	"log"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/handlers/notes"
	"github.com/aflashyrhetoric/lantern-go/handlers/person"

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
	r.GET("/people", person.GetPeople)
	r.POST("/people", person.CreatePerson)
	r.GET("/people/:id", person.ShowPerson)
	r.PUT("/people/:id", person.UpdatePerson)
	r.DELETE("/people/:id", person.DeletePerson)

	// Notes
	r.GET("/notes", notes.GetAllNotes)
	r.POST("/notes", notes.CreateNote)
	r.PUT("/notes/:id", notes.UpdateNote)
	r.GET("/notes/:id", notes.ShowNote)
	r.DELETE("/notes/:id", notes.DeleteNote)

	// Run
	r.Run() // listen and serve on localhost:8080
}
