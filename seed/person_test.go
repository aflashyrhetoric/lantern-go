package seed

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/bxcodec/faker/v3"
)

func randomFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func randomCareer() string {
	return randomFromList([]string{"developer", "slp", "teacher", "singer"})
}

func TestSeedPeople(t *testing.T) {

	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	limit := 40

	for i := 0; i < limit; i++ {

		a := db.Person{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Career:    randomFromList([]string{"developer", "slp", "teacher", "singer"}),
			Mobile:    faker.Phonenumber(),
			Email:     faker.Email(),
			Address:   randomFromList([]string{"1 Infinite Loop, Cupertino CA 11234", "15 Richardson Blvd, Aslip NY 11245"}),
		}

		birthday, err := time.Parse("2006-01-02", faker.Date())
		if err != nil {
			panic(err)
		}
		a.DOB = birthday

		if err != nil {
			panic(err)
		}

		err = db.CreatePerson(a)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%d printed", limit)

}
