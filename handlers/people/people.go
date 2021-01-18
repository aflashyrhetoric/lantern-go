package people

import (
	"time"

	"github.com/gin-gonic/gin"
)

// People .. Represents a Person record in the dossier
type People struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Career    string    `json:"career"`
	Mobile    int       `json:"mobile"`
	Email     int       `json:"email"`
	Address   int       `json:"address"`
	DOB       time.Time `json:"dob"`
}

func GetPeople(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
