package mine

import (
	"unicode"
)

func Punc(str string) string {
	sRunes := []rune(str)
	var words []rune
	for i, values := range sRunes {
		if values != ' ' && i < len(sRunes) && !(unicode.IsDigit(values)) && !(unicode.IsLetter(values)) && values != '\'' {
			if sRunes[i-1] == ' ' {
				sRunes[i], sRunes[i-1] = sRunes[i-1], sRunes[i] 
			}
		}
	}
	for i, chr := range sRunes {
		if i < len(sRunes)-1 && chr == ' ' && sRunes[i-1] == ' ' {
			continue
		}
		words = append(words, sRunes[i])
	}
	return string(words)
}
