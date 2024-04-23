package mine

import (
	"strconv"
	"strings"
	"fmt"
)

func Cap(str string) string {
	content := strings.Split(str, " ")
	for i := 0; i < len(content); i++ {
		if strings.Contains(content[i], "(cap") {
			if strings.Contains(content[i], "(cap,") {
				numb, err:= strconv.Atoi(content[i+1][:len(content[i+1])-1])
				if err != nil || numb > i {
					fmt.Printf("Error at conversion '%v' or cap '%v' is out of range.", err, numb)
				}
				for j := i - numb; j < i; j++ {
					content[j] = strings.ToUpper(string(content[j][0])) + strings.ToLower(string(content[j][1:]))
				}
				content = append(content[:i], content[i+2:]...)
			} else {
				content[i-1] = strings.ToUpper(string(content[i-1][0])) + strings.ToLower(string(content[i-1][1:]))
				content = append(content[:i], content[i+1:]...)
			}
		}
	}
	return strings.Join(content, " ")
}
