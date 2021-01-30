package seed

import (
	"fmt"
	"log"
	"testing"

	"github.com/aflashyrhetoric/lantern-go/db"
	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/aflashyrhetoric/lantern-go/utils"
	"github.com/bxcodec/faker/v3"
)

var PressurePointCount int = 15

func TestSeedPressurePoints(t *testing.T) {

	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < PressurePointCount; i++ {

		a := db.PressurePoint{
			PressurePoint: &models.PressurePoint{
				Description: faker.Sentence(),
				PersonID:    utils.RandomNumNonZero(PersonCount),
			},
		}
		if err != nil {
			panic(err)
		}

		err = db.CreatePressurePoint(&a)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%d printed", PressurePointCount)

}
