package syllable

import (
	"math/rand"
	"time"
)

var vowels = "aeiouy"
var consonants = "bcdfgjklmnpqstvxzhrwy"

func New() string {
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(len(consonants))
	v := rand.Intn(len(vowels))
	consonant := consonants[c]
	vowel := vowels[v]

	s := []byte{consonant, vowel}

	return string(s)
}
