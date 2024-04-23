package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		// If the current word is "(hex)" and it is not the first word in the string.
		if words[i] == "(hex)" && i > 0 {
			// Convert the preceding word (hexadecimal number) to decimal.
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			// Replace the hexadecimal number with its decimal equivalent.
			words[i-1] = fmt.Sprint(num)
			// Remove the "(hex)" marker from the words slice.
			words = append(words[:i], words[i+1:]...)
		}
	}
	// Join the modified words back into a string with spaces.
	return strings.Join(words, " ")
}
