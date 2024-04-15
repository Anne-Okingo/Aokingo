package mine

import "strings"

func AToAn(str string) string {
	// to split string str into words
	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		// change to lowwercase for easy comparison
		if strings.ToLower(words[i]) == "a" {
			//    to get to the first letter in the next word
			if i+1 < len(words) {
				firstletter := strings.ToLower(string(words[i+1][0]))
				vowels := []string{"a", "e", "i", "o", "u", "h"}
				for _, vowel := range vowels {
					if firstletter == vowel {
						words[i] = "an"
						break
					}
				}
			}
		}
	}
	// join the modified words back into one string
	return strings.Join(words, " ")
}
