package mine

import (
	"strconv"
	"strings"
)

func Up(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up") {
			if strings.Contains(words[i], "(up,") {
				number, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i - number; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}
