package mine

import (
	"strconv"
	"strings"
)

func Cap(str string) string {
	content := strings.Split(str, " ")
	for i := 0; i < len(content); i++ {
		if strings.Contains(content[i], "(cap") {
			if strings.Contains(content[i], "(cap,") {
				numb, _ := strconv.Atoi(content[i+1][:len(content[i+1])-1])
				for j := i - numb; j < i; j++ {
					content[j] = strings.Title(content[j])
				}
				content = append(content[:i], content[i+2:]...)
			} else {
				content[i-1] = strings.Title(content[i-1])
				content = append(content[:i], content[i+1:]...)
			}
		}
	}
	return strings.Join(content, " ")
}
