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

// var JWTSigningKey string

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	// config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:3000", "https://lantern.vercel.app", "https://staging-lantern.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type, *"}
	r.Use(cors.New(config))

	// Add some middleware for cors
	protected := r.Group("/api/")
	// protected.Use(cors.New(config))

	// Load ENV Variables
	ENV := os.Getenv("LANTERN_ENV")
	var connectionString string

	if ENV == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		DBHOST := os.Getenv("DB_HOST")
		DBPORT := os.Getenv("DB_PORT")
		DBUSER := os.Getenv("DB_USER")
		DBDATABASE := os.Getenv("DB_DATABASE")
		DBSSLMODE := os.Getenv("DB_SSLMODE")
		// JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")

		port, _ := strconv.Atoi(DBPORT)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", DBHOST, port, DBUSER, DBDATABASE, DBSSLMODE)
	} else {

		DBHOST := os.Getenv("DB_HOST")
		DBPORT := os.Getenv("DB_PORT")
		DBUSER := os.Getenv("DB_USER")
		DBPASSWORD := os.Getenv("DB_PASSWORD")
		DBDATABASE := os.Getenv("DB_DATABASE")
		DBSSLMODE := os.Getenv("DB_SSLMODE")

		port, _ := strconv.Atoi(DBPORT)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", DBHOST, port, DBUSER, DBPASSWORD, DBDATABASE, DBSSLMODE)
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
