package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = fmt.Sprint(num)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
