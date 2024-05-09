package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Up(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up") || strings.Contains(words[i], "(UP") || strings.Contains(words[i], "(uP") || strings.Contains(words[i], "(Up") {
			if strings.Contains(words[i], "(up,") || strings.Contains(words[i], "(UP,") || strings.Contains(words[i], "(uP,") || strings.Contains(words[i], "(Up,") {
				num, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				if err != nil {
					fmt.Println("error in conversion to integr", err)
				}
				if num <= len(words[:i]) {
					for j := i - num; j < i; j++ {
						words[j] = strings.ToUpper(string(words[j]))
					}
					words = append(words[:i], words[i+2:]...)
				} else if num >= len(words[:i]) {
					for r := 0; r < len(words[:i]); r++ {
						words[r] = strings.ToUpper(string(words[r]))
					}
					words = append(words[:i], words[i+2:]...)
				}
			} else {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}
