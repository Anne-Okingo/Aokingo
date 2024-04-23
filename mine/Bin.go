package mine

import(
	"strings"
	"strconv"
	"fmt"

)

func Bin(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++{
		// If the current word is "(bin)" and it is not the first word in the string.
		if words[i] == "(bin)"{
			// Convert the preceding word (binary number) to decimal.
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			// Replace the binary number with its decimal equivalent.
			words[i-1] = fmt.Sprint(num)
			// Remove the "(bin)" marker from the words slice.
			words = append(words[:i], words[i+1:]...)
		}
	}

	// Join the modified words back into a string with spaces.
	return strings.Join(words, " ")
}