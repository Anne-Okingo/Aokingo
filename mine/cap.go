package mine

import (
	"strconv"
	"strings"
)

func Cap(str string) string {
	content := strings.Fields(str)
	for j := 0; j < len(content); j++ {
		if strings.Contains(content[j], "(cap") {
			if strings.Contains(content[j], "(cap,") {
				num, _ := strconv.Atoi(content[j+1][:len(content[j+1])-1])
				//for i := j-num; i < j; i++ 
				for i := j-1; i > j-num; i-- {
					content[i] = strings.Title(content[i])
				}
				content = append(content[:j], content[j+2:]...)
			} else {
				content[j-1] = strings.Title(content[j-1])
				content = append(content[:j], content[j+1:]...)
			}
		}
	}
	return strings.Join(content, " ")
}
