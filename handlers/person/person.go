package person

import (
	"net/http"

	"github.com/aflashyrhetoric/lantern-api/db"
	"github.com/gin-gonic/gin"
)

// GetPeople .. Returns all People
func GetPeople(c *gin.Context) {
	people, err := db.GetAllPeople()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"data": people,
	})
}
