package faker

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/db"
	faker "github.com/bxcodec/faker/v3"
)

func PersonTestWithTags() {
	a := db.Person{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
}
