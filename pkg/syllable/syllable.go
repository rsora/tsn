// syllable is a package that generates a random syllable.
// A syllable fot tsn is a 2 char string that starts with a consonant
// end ends with a vowel.
package syllable

import (
	"math/rand"
	"time"
)

var vowels = "aeiouy"
var consonants = "bcdfgjklmnpqstvxzhrwy"

// New returns a syllable string generated randomly.
func New() string {
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(len(consonants))
	v := rand.Intn(len(vowels))
	consonant := consonants[c]
	vowel := vowels[v]

	s := []byte{consonant, vowel}

	return string(s)
}
