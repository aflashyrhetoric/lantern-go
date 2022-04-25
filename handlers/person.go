package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

type PersonPageResponse struct {
	Person   *models.Person  `json:"person"`
	UserData models.UserData `json:"user_data"`
}

// GetPeople .. Returns all People
func GetPeople(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user_id is not valid in context"))
	}
	// userID := fmt.Sprint(userIDInterface)
	userID := userIDInterface.(int64)

	people, err := db.GetAllPeople(userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func ShowPerson(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user_id is not valid in context"))
	}

	// userID := fmt.Sprint(userIDInterface)
	userID := userIDInterface.(int64)

	id := c.Param("id")
	person, err := db.GetPerson(id, userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if person.UserID != userID {
		c.AbortWithError(http.StatusForbidden, fmt.Errorf("user %d tried to access user %d to which they do not have permission - go away", person.UserID, userID))
	}

	people, err := db.GetAllPeople(userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": PersonPageResponse{
			Person: person,
			UserData: models.UserData{
				People: people,
			},
		},
	})
}

func CreatePerson(c *gin.Context) {
	var (
		dob time.Time
		err error
	)

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user_id is not valid in context"))
	}
	userID := userIDInterface.(int64)

	dbModel := &models.CreatePersonRequest{}
	err = c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	person := db.Person{
		Person: &models.Person{
			FirstName:          dbModel.FirstName,
			LastName:           dbModel.LastName,
			Career:             dbModel.Career,
			Mobile:             dbModel.Mobile,
			Email:              dbModel.Email,
			Address:            dbModel.Address,
			UserID:             userID,
			RelationshipToUser: dbModel.RelationshipToUser,
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
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not find user with id %s", id))
	}

	updatePersonReq := &models.UpdatePersonRequest{}
	err = c.BindJSON(&updatePersonReq)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not bind the values to JSON for person %v", err))
	}

	if updatePersonReq.FirstName != "" {
		person.FirstName = updatePersonReq.FirstName
	}
	if updatePersonReq.LastName != "" {
		person.LastName = updatePersonReq.LastName
	}
	// if updatePersonReq.Career != "" {
	person.Career = updatePersonReq.Career
	// }
	// if updatePersonReq.Mobile != "" {
	person.Mobile = updatePersonReq.Mobile
	// }
	// if updatePersonReq.Email != "" {
	person.Email = updatePersonReq.Email
	// }
	// if updatePersonReq.Address != "" {
	person.Address = updatePersonReq.Address
	// }

	person.RelationshipToUser = updatePersonReq.RelationshipToUser

	if updatePersonReq.RelationshipToUserThroughPerson != 0 && updatePersonReq.RelationshipToUser != "" {
		person.RelationshipToUserThroughPerson = updatePersonReq.RelationshipToUserThroughPerson
	}

	var dob time.Time
	if updatePersonReq.DOB != nil {
		dob, err = time.Parse("2006-01-02", *updatePersonReq.DOB)
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
