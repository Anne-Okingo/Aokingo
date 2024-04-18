package mine

import "unicode"

func Punc(str string) string {
	runes := []rune(str)
	var words []rune
	for i := 0; i < len(runes); i++ {
		if runes[i] != ' ' && i < len(runes) && !(unicode.IsDigit(runes[i])) && !(unicode.IsLetter(runes[i])) && runes[i] != '\'' {
			if runes[i-1] == ' ' {
				runes[i], runes[i-1] = runes[i-1], runes[i]
			}
		}
	}
	for i, chr := range runes {
	if i < len(runes) && chr == ' ' && runes[i-1] == ' ' {
			continue
		}
		words = append(words, runes[i])
	}
	return string(words)
}