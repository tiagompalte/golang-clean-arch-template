package test

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Email() string {
	email := uuid.NewString()
	return fmt.Sprintf("%s@email.com", email)
}

func RandomName() string {
	charSet := []rune("abcdefghijklmnopqrstuvwxyz")
	var output strings.Builder

	for i := 0; i < 10; i++ {
		output.WriteRune(charSet[rand.Intn(len(charSet))])
	}
	output.WriteRune(' ')
	for i := 0; i < 10; i++ {
		output.WriteRune(charSet[rand.Intn(len(charSet))])
	}

	caser := cases.Title(language.BrazilianPortuguese)
	name := caser.String(output.String())

	return name
}
