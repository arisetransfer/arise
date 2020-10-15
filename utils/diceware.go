package utils

import (
	"log"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
)

func Dice(i int) string {
	list, err := diceware.Generate(i)
	if err != nil {
		log.Fatal(err)
	}
	code := strings.Join(list, "-")
	return code
}
