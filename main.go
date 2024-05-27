package main

import (
	"fmt"
	"github.com/eduardoraider/go-fuzzy-test/words"
)

func main() {
	word := "Eduardo"

	reversedWord, err := words.Reverse(word)
	if err != nil {
		panic(err)
	}

	fmt.Println(reversedWord)
}
