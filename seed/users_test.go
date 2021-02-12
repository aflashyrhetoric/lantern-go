package seed

import (
	"fmt"
	"log"
	"testing"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/bxcodec/faker/v3"
)

var UserCount int = 20

func TestSeedUsers(t *testing.T) {

	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < UserCount; i++ {

		a := db.User{
			User: &models.User{
				Email:    faker.Username(),
				Password: faker.Password(),
			},
		}
		if err != nil {
			panic(err)
		}

		err = db.CreateUser(&a)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%d printed", UserCount)

}
