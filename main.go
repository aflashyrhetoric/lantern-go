package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/handlers"
	"github.com/joho/godotenv"

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

	// Load ENV Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_ENV := os.Getenv("ENV")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	var connectionString string
	port, _ := strconv.Atoi(DB_PORT)

	if DB_ENV == "development" {
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", DB_HOST, port, DB_USER, DB_DATABASE, DB_SSLMODE)
	} else {
		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", DB_HOST, port, DB_USER, DB_PASSWORD, DB_DATABASE, DB_SSLMODE)
	}

	fmt.Print(connectionString)

	// Use the InitDB function to initialise the global variable.
	err = db.Start(connectionString)
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
