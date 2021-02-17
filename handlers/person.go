package handlers

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
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user_id is not valid in context"))
	}
	userID := fmt.Sprint(userIDInterface)

	people, err := db.GetAllPeople(userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func ShowPerson(c *gin.Context) {
	id := c.Param("id")
	person, err := db.GetPerson(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &person,
	})
}

func CreatePerson(c *gin.Context) {
	var (
		dob time.Time
		err error
	)

	dbModel := &models.PersonRequest{}
	err = c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	person := db.Person{
		Person: &models.Person{
			FirstName: dbModel.FirstName,
			LastName:  dbModel.LastName,
			Career:    dbModel.Career,
			Mobile:    dbModel.Mobile,
			Email:     dbModel.Email,
			Address:   dbModel.Address,
		},
	}

	if dbModel.DOB != nil {
		dob, err = time.Parse("2006-01-02", *dbModel.DOB)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not parse dob value from the post message"))
		}

		if dob.IsZero() {
			person.DOB = nil
		} else {
			person.DOB = &dob
		}
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

	person, err := db.GetPersonalData(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Could not find user with id %s", id))
	}

	dbModel := &models.PersonRequest{}
	err = c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not bind the values to JSON for person %v", err))
	}

	if dbModel.FirstName != "" {
		person.FirstName = dbModel.FirstName
	}
	if dbModel.LastName != "" {
		person.LastName = dbModel.LastName
	}
	if dbModel.Career != "" {
		person.Career = dbModel.Career
	}
	if dbModel.Mobile != "" {
		person.Mobile = dbModel.Mobile
	}
	if dbModel.Email != "" {
		person.Email = dbModel.Email
	}
	if dbModel.Address != "" {
		person.Address = dbModel.Address
	}

	var dob time.Time
	if dbModel.DOB != nil {
		dob, err = time.Parse("2006-01-02", *dbModel.DOB)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not parse dob value from the post message"))
		}
	}

	if dob.IsZero() {
		person.DOB = nil
	} else {
		person.DOB = &dob
	}

	// spew.Dump(person)

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
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the DeletePerson func was not provided"))
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
