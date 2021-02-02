package handlers

import (
	"fmt"
	"net/http"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

// GetUsers .. Returns all Users

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	user, err := db.GetUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

func CreateUser(c *gin.Context) {
	dbModel := &models.UserRequest{}
	err := c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	user := db.User{
		User: &models.User{
			Email:    dbModel.Email,
			Password: dbModel.Password,
		},
	}

	err = db.CreateUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

// func UpdateUser(c *gin.Context) {
// 	if c.Param("id") == "" {
// 		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the UpdateUser func was not provided"))
// 	}
// 	var id = c.Param("id")

// 	user, err := db.GetUserData(id)
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Could not find user with id %s", id))
// 	}

// 	dbModel := &models.UserRequest{}
// 	err = c.BindJSON(&dbModel)
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not bind the values to JSON for  %v", err))
// 	}

// 	if dbModel.Email != "" {
// 		user.Email = dbModel.Email
// 	}

// 	if dbModel.Password != "" {
// 		user.Email = dbModel.Email
// 	}

// 	err = db.UpdateUser(id, user)
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while updating user %s with error: %v", id, err))
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": user,
// 	})
// }

func DeleteUser(c *gin.Context) {
	if c.Param("id") == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("ID Parameter to the DeleteUser func was not provided"))
	}
	var id = c.Param("id")

	err := db.DeleteUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error while deleting user %s with error: %v", id, err))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("user with id %s was successfully deleted", id),
	})
}
