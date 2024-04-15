// package mine

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func Bin(str string) string {
// 	words := strings.Split(str, " ")
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(bin)" {
// 			num, _ := strconv.ParseInt(words[i-1], 2, 64)
// 			words[i-1] = fmt.Sprint(num)
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
package mine

import(
	"strings"
	"strconv"
	"fmt"

)

func Bin(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++{
		if words[i] == "(bin)"{
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = fmt.Sprint(num)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}