package main

import (
	"fmt"
	"math/rand"
)

func GenerateRandomAlnumString(args ...int) string {
	upperBound := 20
	lowerBound := 5
	if len(args) >= 1 {
		upperBound = args[0]
	}
	if len(args) >= 2 {
		lowerBound = args[1]
	}
	length := rand.Intn(upperBound-lowerBound) + lowerBound
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = byte(GenerateRandomAlnumCharacter())
	}
	return string(str)
}

func GenerateRandomAlnumStringN(length int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = byte(GenerateRandomAlnumCharacter())
	}
	return string(str)
}

func GenerateRandomAlnumCharacter() byte {
	return ALNUM[rand.Int()%len(ALNUM)]
}

func GenerateRandomIntInRange(rangeFloor int, rangeCeiling int) (int, error) {
	if rangeCeiling <= rangeFloor {
		return -1, fmt.Errorf("error: GenerateRandomIntInRange must be provided with a larger ceiling than the floor")
	}
	return rand.Int()%(rangeCeiling-rangeFloor) + rangeFloor, nil
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}
