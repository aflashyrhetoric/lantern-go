package person

import (
	"net/http"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
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

func CreatePerson(c *gin.Context) {
	birthday, err := time.Parse(time.RFC3339, c.PostForm("dob"))

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	person := db.Person{
		FirstName: c.PostForm("first_name"),
		LastName:  c.PostForm("last_name"),
		Career:    c.PostForm("career"),
		Mobile:    c.PostForm("mobile"),
		Email:     c.PostForm("email"),
		Address:   c.PostForm("address"),
		DOB:       birthday,
	}

	err = db.CreatePerson(person)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"data": person,
	})
}
