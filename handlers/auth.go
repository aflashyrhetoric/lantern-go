package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SIGNING_KEY string

func SignupUser(c *gin.Context) {
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
			CreatedAt: now,
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
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not parse the user login request"))
	}

	user, err := db.GetUserByEmail(dbModel.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not find user by email"))
	}

	userPW := user.Password
	reqPW := dbModel.Password

	err = bcrypt.CompareHashAndPassword([]byte(userPW), []byte(reqPW))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("password does not match"))
	}

	// Create the token
	JWT_SIGNING_KEY = os.Getenv("JWT_SIGNING_KEY")
	if JWT_SIGNING_KEY == "" {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("JWT_SIGNING_KEY is not set"))
	}
	mySigningKey := []byte(JWT_SIGNING_KEY)
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 5).Unix(),
	}
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not get the signed string from the token"))
	}

	ENV := os.Getenv("LANTERN_ENV")

	if ENV == "development" {
		c.SetCookie("authorized_user", tokenString, int(time.Second)*60*24*5, "/", "", false, true)
	} else {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("authorized_user", tokenString, int(time.Second)*60*24*5, "/", "", true, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokenString,
	})
}

func Logout(c *gin.Context) {
	ENV := os.Getenv("LANTERN_ENV")

	t := time.Now()
	t2 := t.AddDate(-1, 0, 0)
	t3 := int(t2.Unix())

	if ENV == "development" {
		c.SetCookie("authorized_user", "nil", t3, "/", "", false, true)
	} else {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("authorized_user", "nil", t3, "/", "", true, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Logged out",
	})
}
