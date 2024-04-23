package mine

import "unicode"

func Punc(str string) string {
	runes := []rune(str)
	// Initialize a slice to store modified runes.
	var words []rune
	for i := 0; i < len(runes); i++ {
		// If the current rune is not a space and not a digit, letter, or apostrophe, and it's not the last character in the string.
		if runes[i] != ' ' && i < len(runes) && !(unicode.IsDigit(runes[i])) && !(unicode.IsLetter(runes[i])) && runes[i] != '\'' {
			// If the preceding character is a space, swap the current and preceding runes.
			if runes[i-1] == ' ' {
				runes[i], runes[i-1] = runes[i-1], runes[i]
			}
		}
	}
	// Iterate over the modified runes to remove duplicate spaces.
	for i, chr := range runes {
		// If the current rune is a space and the preceding rune is also a space, continue to the next iteration.
	if i < len(runes) && chr == ' ' && runes[i-1] == ' ' {
			continue
		}
		// Append non-duplicate runes to the words slice.
		words = append(words, runes[i])
	}
	// Convert the modified runes back to a string.
	return string(words)
}