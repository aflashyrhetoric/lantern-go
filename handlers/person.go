package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AppledoreClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GetPeople .. Returns all People
func GetPeople(c *gin.Context) {
	cookie, err := c.Request.Cookie("authorized_user")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// For any other type of error, return a bad request status
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tokenString := cookie.Value

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, &AppledoreClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretString := os.Getenv("JWT_SIGNING_KEY")
		return []byte(secretString), nil

	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if !tkn.Valid {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var claims *AppledoreClaims
	if claims, ok := tkn.Claims.(*AppledoreClaims); ok && tkn.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

	spew.Dump(claims)

	// people, err := db.GetAllPeople(fmt.Sprint(claims.UserID))
	people, err := db.GetAllPeople("3")
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
