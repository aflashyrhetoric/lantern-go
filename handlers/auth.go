package handlers

import (
	"net/http"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	dbModel := &models.UserRequest{}
	err := c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	now := time.Now()

	user := db.User{
		User: &models.User{
			Email:     dbModel.Email,
			Password:  dbModel.Password,
			CreatedAt: &now,
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

func LoginUser(c *gin.Context) {
	dbModel := &models.UserRequest{}
	err := c.BindJSON(&dbModel)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	user, err := db.GetUserByEmail(dbModel.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// token := jwt.New(jwt.SigningMethodRS512)
	// claims := make(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix()
	// claims["iat"] = time.Now().Unix()
	// token.Claims = claims

	// BY HERE: User is created
	c.JSON(http.StatusCreated, gin.H{
		"data": user.Password,
	})
}
