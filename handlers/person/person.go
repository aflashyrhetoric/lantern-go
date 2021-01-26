package person

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

// GetPeople .. Returns all People
func GetPeople(c *gin.Context) {
	people, err := db.GetAllPeople()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func ShowPerson(c *gin.Context) {
	id := c.Param("id")
	person, err := db.GetPersonWithID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &person,
	})
}

func CreatePerson(c *gin.Context) {
	birthday, err := time.Parse("2006-01-02", c.PostForm("dob"))

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	person := db.Person{
		Person: &models.Person{
			FirstName: c.PostForm("first_name"),
			LastName:  c.PostForm("last_name"),
			Career:    c.PostForm("career"),
			Mobile:    c.PostForm("mobile"),
			Email:     c.PostForm("email"),
			Address:   c.PostForm("address"),
			DOB:       &birthday,
		},
	}

	err = db.CreatePerson(&person)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": person,
	})
}

func UpdatePerson(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the UpdatePerson func was not provided"))
	}
	var id = c.Param("id")

	person, err := db.GetPersonWithID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Could not find user with id %s", id))

	}

	if c.PostForm("first_name") != "" {
		person.FirstName = c.PostForm("first_name")
	}
	if c.PostForm("last_name") != "" {
		person.LastName = c.PostForm("last_name")
	}
	if c.PostForm("career") != "" {
		person.Career = c.PostForm("career")
	}
	if c.PostForm("mobile") != "" {
		person.Mobile = c.PostForm("mobile")
	}
	if c.PostForm("email") != "" {
		person.Email = c.PostForm("email")
	}
	if c.PostForm("address") != "" {
		person.Address = c.PostForm("address")
	}

	var dob time.Time

	if c.PostForm("dob") != "" {
		dob, err = time.Parse(time.RFC3339, c.PostForm("dob"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not parse dob value from the post message"))
		}

		if dob.IsZero() {
			person.DOB = &dob
		} else {
			person.DOB = nil
		}
	}

	err = db.UpdatePerson(id, person)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while updating person %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": person,
	})
}

func DeletePerson(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the UpdatePerson func was not provided"))
	}
	var id = c.Param("id")

	err := db.DeletePerson(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while deleting person %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("person with id %s was successfully deleted", id),
	})
}
