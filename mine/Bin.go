package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Bin(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		// considering all variations on the word BIN
		if (words[i] == "(bin)" ||  words[i] == "(Bin)" ||  words[i] == "(BIN)" ||  words[i] == "(bIN)" ||  words[i] == "(biN)" ||  words[i] == "(BIn)" ||  words[i] == "(BiN)") && i > 0{
			// converting the number is 1E from decimal to binary
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("err in convertion of hexadecimal", err)
			}
			words[i-1] = fmt.Sprint(num)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
