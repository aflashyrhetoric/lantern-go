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

var JWT_SIGNING_KEY string

func main() {
	r := gin.Default()

	protected := r.Group("/")

	// Add some middleware for cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	r.Use(cors.New(config))

	// Load ENV Variables
	ENV := os.Getenv("LANTERN_ENV")
	var connectionString string

	if ENV == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		DB_HOST := os.Getenv("DB_HOST")
		DB_PORT := os.Getenv("DB_PORT")
		DB_USER := os.Getenv("DB_USER")
		DB_DATABASE := os.Getenv("DB_DATABASE")
		DB_SSLMODE := os.Getenv("DB_SSLMODE")
		JWT_SIGNING_KEY = os.Getenv("JWT_SIGNING_KEY")

		port, _ := strconv.Atoi(DB_PORT)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", DB_HOST, port, DB_USER, DB_DATABASE, DB_SSLMODE)
	} else {

		DB_HOST := os.Getenv("DB_HOST")
		DB_PORT := os.Getenv("DB_PORT")
		DB_USER := os.Getenv("DB_USER")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_DATABASE := os.Getenv("DB_DATABASE")
		DB_SSLMODE := os.Getenv("DB_SSLMODE")

		port, _ := strconv.Atoi(DB_PORT)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", DB_HOST, port, DB_USER, DB_PASSWORD, DB_DATABASE, DB_SSLMODE)
	}

	// Use the InitDB function to initialise the global variable.
	err := db.Start(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Auth
	protected.POST("/auth/signup", handlers.SignupUser)
	protected.POST("/auth/login", handlers.LoginUser)

	// Routes
	protected.GET("/people", handlers.GetPeople)
	protected.POST("/people", handlers.CreatePerson)
	protected.GET("/people/:id", handlers.ShowPerson)
	protected.PUT("/people/:id", handlers.UpdatePerson)
	protected.DELETE("/people/:id", handlers.DeletePerson)

	// Notes
	protected.POST("/notes/:person_id", handlers.CreateNote)
	protected.DELETE("/notes/:id", handlers.DeleteNote)

	// Pressure Points
	protected.POST("/pressure-points/:person_id", handlers.CreatePressurePoint)
	protected.DELETE("/pressure-points/:id", handlers.DeletePressurePoint)

	// Users
	// r.POST("/users", handlers.CreateUser)
	// r.DELETE("/notes/:id", handlers.DeleteNote)

	// Run
	r.Run() // listen and serve on localhost:8080
}
