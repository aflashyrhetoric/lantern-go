package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"net/http"
	_ "net/http"
)

type AppClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("authorized_user")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.Next()
				// c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			// For any other type of error, return a bad request status
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		tokenString := cookie.Value

		if tokenString == "nil" {
			c.Next()
			return
		}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			secretString := os.Getenv("JWT_SIGNING_KEY")
			return []byte(secretString), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if !tkn.Valid {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claims := tkn.Claims.(*AppClaims)
		c.Set("user_id", claims.UserID)
	}
}

var ENV string

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	// config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:3000", "https://lantern.vercel.app", "https://staging-lantern.vercel.app/appledore"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type, *"}
	r.Use(cors.New(config))

	// Add some middleware for cors
	protected := r.Group("/api/")
	protected.Use(Authenticate())
	// protected.Use(cors.New(config))

	// Load ENV Variables
	ENV = os.Getenv("LANTERN_ENV")
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
	r.POST("/api/auth/signup", handlers.SignupUser)
	r.POST("/api/auth/login", handlers.LoginUser)
	r.POST("/api/auth/logout", handlers.Logout) // Keep as POST, so that browser pre-fetching does not invalidate our sesssion

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
