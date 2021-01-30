package utils

import (
	"math/rand"
)

func RandomNumNonZero(max int) int {
	v := rand.Intn(max)
	if v == 0 {
		return 1
	}

	return v
}

func RandomNum(max int) int {
	return rand.Intn(max)
}

func RandomStringFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func RandomIntFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func RandomFromList(list []interface{}) interface{} {
	return list[rand.Intn(len(list))]
}
