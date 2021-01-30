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

var NoteCount int = 20

func TestSeedNotes(t *testing.T) {

	err := db.Start("user=ko dbname=lantern-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < NoteCount; i++ {

		a := db.Note{
			Note: &models.Note{
				Text:     faker.Sentence(),
				PersonID: utils.RandomNumNonZero(PersonCount),
			},
		}
		if err != nil {
			panic(err)
		}

		err = db.CreateNote(&a)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%d printed", NoteCount)

}
