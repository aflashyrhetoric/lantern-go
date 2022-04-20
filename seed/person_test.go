package seed

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/aflashyrhetoric/lantern-go/utils"
	"github.com/bxcodec/faker/v3"
)

func randomCareer() string {
	return utils.RandomStringFromList([]string{"developer", "slp", "teacher", "singer"})
}

var PersonCount int = 21

func TestSeedPeople(t *testing.T) {

	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < PersonCount; i++ {

		a := db.Person{
			Person: &models.Person{
				FirstName: faker.FirstName(),
				LastName:  faker.LastName(),
				Career:    utils.RandomStringFromList([]string{"developer", "slp", "teacher", "singer"}),
				Mobile:    faker.Phonenumber(),
				Email:     faker.Email(),
				Address:   utils.RandomStringFromList([]string{"1 Infinite Loop, Cupertino CA 11234", "15 Richardson Blvd, Aslip NY 11245"}),
			},
		}

		birthday, err := time.Parse("2006-01-02", faker.Date())
		if err != nil {
			panic(err)
		}
		a.DOB = &birthday
		a.UserID = int64(utils.RandomNum(PersonCount))

		if err != nil {
			panic(err)
		}

		err = db.CreatePerson(&a)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%d printed", PersonCount)
}
