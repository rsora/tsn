package main

import (
	"fmt"

	"github.com/rsora/tsn/internal/syllable"
)

const syllables = 2

func main() {
	var s = ""
	for i := 0; i < syllables; i++ {
		s = fmt.Sprint(s, syllable.New())
	}

	fmt.Println(s)
}
