package mine

import (
	"strings"
)


func AToAn(str string) string {
	// Split the input string into words.
	words := strings.Fields(str)
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(words); i++ {
		for _, firstL := range vowels {
			if words[i] == "a" && string(words[i+1][0]) == firstL {
				words[i] = "an"
			} else {
				// If the current word is 'a' or 'A' and the next word starts with a vowel or 'h', replace 'a'/'A' with 'an'/'An'
				if words[i] == "A" && string(words[i+1][0]) == firstL {
					words[i] = "An"
				}
			}
		}
	}
	// Join the modified words back into a string with spaces.
	return strings.Join(words, " ")
}
