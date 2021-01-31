package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

// CreatePressurePoint .. creates a new pressure point
func CreatePressurePoint(c *gin.Context) {
	personIDParam := c.Param("person_id")
	personID, err := strconv.Atoi(personIDParam)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	pressurepointRequest := models.PressurePointRequest{}
	err = c.BindJSON(&pressurepointRequest)

	pressurepoint := db.PressurePoint{
		PressurePoint: &models.PressurePoint{
			Description: pressurepointRequest.Description,
			PersonID:    personID,
		},
	}

	err = db.CreatePressurePoint(&pressurepoint)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": pressurepoint,
	})
}

// DeletePressurePoint .. deletes a pressure point
func DeletePressurePoint(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the DeletePressurePoint func was not provided"))
	}
	var id = c.Param("id")

	err := db.DeletePressurePoint(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while deleting pressurepoint %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("pressurepoint with id %s was successfully deleted", id),
	})
}
