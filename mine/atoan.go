package mine

import (
	"strings"
)

func AToAn(str string) string {
	words := strings.Fields(str)
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(words); i++ {
		for _, firstL := range vowels {
			if words[i] == "a" && string(words[i+1][0]) == firstL {
				words[i] = "an"
			} else {
				if words[i] == "A" && string(words[i+1][0]) == firstL {
					words[i] = "An"
				}
			}
		}
	}
	return strings.Join(words, " ")
}
