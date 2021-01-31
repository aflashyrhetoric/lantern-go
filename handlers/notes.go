package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

// CreateNote .. creates a new note
func CreateNote(c *gin.Context) {
	personIDParam := c.Param("person_id")
	personID, err := strconv.Atoi(personIDParam)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	noteRequest := models.NoteRequest{}
	err = c.BindJSON(&noteRequest)

	note := db.Note{
		Note: &models.Note{
			Text:     noteRequest.Text,
			PersonID: personID,
		},
	}

	err = db.CreateNote(&note)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": note,
	})
}

// DeleteNote .. deletes a new note
func DeleteNote(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the DeleteNote func was not provided"))
	}
	var id = c.Param("id")

	err := db.DeleteNote(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while deleting note %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("note with id %s was successfully deleted", id),
	})
}
