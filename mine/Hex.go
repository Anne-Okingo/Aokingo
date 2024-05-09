package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		// considering all variations on the word BIN
		if (words[i] == "(hex)" ||  words[i] == "(Hex)" ||  words[i] == "(HEX)" ||  words[i] == "(hEX)" ||  words[i] == "(heX)" ||  words[i] == "(HEx)" ||  words[i] == "(HeX)") && i >= 1 {
			// Convert the preceding word (hexadecimal number) to decimal.
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error in conversion of integer", err)
			}
			// Replace the hexadecimal number with its decimal equivalent.
			words[i-1] = fmt.Sprint(num)
			// Remove the "(hex)" marker from the words slice.
			words = append(words[:i], words[i+1:]...)
		}
	}
	// Join the modified words back into a string with spaces.
	return strings.Join(words, " ")
}
