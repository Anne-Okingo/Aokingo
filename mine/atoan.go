package mine

import (
	"strings"
)

func AToAn(str string) string {
	words := strings.Split(str, " ")

	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(words); i++ {
		for _, firstLetter := range vowels {
			if words[i] == "a" && string(words[i+1][0]) == firstLetter {
				words[i] = "an"
			} else {
				if words[i] == "A" && string(words[i+1][0]) == firstLetter {
					words[i] = "An"
				}
			}
		}
	}
	return strings.Join(words, " ")
}
