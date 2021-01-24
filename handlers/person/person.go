package person

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/davecgh/go-spew/spew"
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

func ShowPerson(c *gin.Context) {
	id := c.Param("id")
	person, err := db.ShowPerson(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"data": &person,
	})
}

func CreatePerson(c *gin.Context) {
	spew.Dump(c)
	birthday, err := time.Parse("2006-01-02", c.PostForm("dob"))

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
		DOB:       &birthday,
	}

	err = db.CreatePerson(&person)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"data": person,
	})
}

func UpdatePerson(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the UpdatePerson func was not provided"))
	}
	var id = c.Param("id")

	person, err := db.ShowPerson(id)
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

	dob, err := time.Parse(time.RFC3339, c.PostForm("dob"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if dob.IsZero() {
		person.DOB = &dob
	} else {
		person.DOB = nil
	}

	err = db.UpdatePerson(id, person)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"data": person,
	})
}
