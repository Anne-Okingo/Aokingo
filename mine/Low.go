package mine

import (
	"strconv"
	"strings"
)

func Low(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(low") {
			if strings.Contains(words[i], "(low,") {
				numb, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i-numb; j < i; j++{//j := i - 1; j > i-numb; i-- {
					words[j] = strings.ToLower(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}
