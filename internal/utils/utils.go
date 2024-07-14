package utils

import "github.com/mazen160/go-random"

func GenerateShortID() (string, error) {
	return random.String(8)
}

func GenerateShortIDWithLength(value int) (string, error) {
	return random.String(value)
}
