package mine

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Low(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(low") || strings.Contains(words[i], "(Low") || strings.Contains(words[i], "(LOW") || strings.Contains(words[i], "(lOW") || strings.Contains(words[i], "(loW") || strings.Contains(words[i], "(lOw") || strings.Contains(words[i], "(LoW"){
			if strings.Contains(words[i], "(low,") || strings.Contains(words[i], "(Low,") || strings.Contains(words[i], "(LOW,") || strings.Contains(words[i], "(lOW,") || strings.Contains(words[i], "(loW,") || strings.Contains(words[i], "(lOw,") || strings.Contains(words[i], "(LoW,"){
				num, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				if err != nil {
					fmt.Println("error in conversion", err)
					os.Exit(0)
				}
				if num <= len(words[:i]) {
					for j := i- num; j < i; j++ {
						words[j] = strings.ToLower(string(words[j]))
					}
					words = append(words[:i], words[i+2:]...)
				} else if num >= len(words[:i]) {
					for r := 0; r < len(words[:i]); r++ {
						words[r] = strings.ToLower(string(words[r]))
					}
					words = append(words[:i], words[i+2:]...)
				} 
			}else {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}