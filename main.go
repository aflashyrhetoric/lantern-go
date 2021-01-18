package main

import (
	"github.com/aflashyrhetoric/lantern-api/handlers/people"
	"github.com/gin-gonic/gin"

	_ "net/http"
)

func main() {
	r := gin.Default()
	r.GET("/pzng", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.GET("/someGet", getting)
	r.GET("/ping", people.Hello)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
