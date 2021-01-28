package notes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

// GetNotes .. Returns all Notes
func GetAllNotes(c *gin.Context) {
	people, err := db.GetAllNotes()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func ShowNote(c *gin.Context) {
	id := c.Param("id")
	note, err := db.GetNoteWithID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &note,
	})
}

func CreateNote(c *gin.Context) {

	person_id := c.PostForm("person_id")
	personID, err := strconv.Atoi(person_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	note := db.Note{
		Note: &models.Note{
			Text:     c.PostForm("text"),
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

func UpdateNote(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the UpdateNote func was not provided"))
	}
	var id = c.Param("id")

	note, err := db.GetNoteWithID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Could not find note with id %s", id))

	}

	if c.PostForm("text") != "" {
		note.Text = c.PostForm("text")
	}

	err = db.UpdateNote(id, note)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while updating note %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": note,
	})
}

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
