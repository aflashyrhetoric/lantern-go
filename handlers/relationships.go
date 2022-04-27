package handlers

import (
	"fmt"
	"net/http"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

func CreateRelationship(c *gin.Context) {
	createRelationshipRequest := &models.CreateRelationshipRequest{}
	err := c.BindJSON(&createRelationshipRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	r := db.Relationship{
		Relationship: &models.Relationship{
			PersonOneID:      createRelationshipRequest.PersonOneID,
			PersonTwoID:      createRelationshipRequest.PersonTwoID,
			RelationshipType: createRelationshipRequest.RelationshipType,
		},
	}

	err = db.CreateRelationship(&r)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": r,
	})
}

func DeleteRelationship(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the DeleteRelationship func was not provided"))
	}
	var id = c.Param("id")

	err := db.DeleteRelationship(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while deleting relationship %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("relationship with id %s was successfully deleted", id),
	})
}
