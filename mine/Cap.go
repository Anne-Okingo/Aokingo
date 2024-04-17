package mine

import (
	"strconv"
	"strings"
)

func Cap(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(cap") {
			if strings.Contains(words[i], "(cap,") {
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i - num; j < i; j++ {
					words[j] = strings.Title(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			}
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
