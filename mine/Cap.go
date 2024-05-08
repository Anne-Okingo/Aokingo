package mine

import (
	"fmt"
	"strconv"
	"strings"
)

func Cap(str string) string {
	// First Split the string using a space
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		// Check if the words contains all variations of the word "(cap", while at it also check for all variations of "(cap,"
		if strings.Contains(words[i], "(cap") || strings.Contains(words[i], "(Cap") || strings.Contains(words[i], "(CAP") || strings.Contains(words[i], "(cAP") || strings.Contains(words[i], "(caP") || strings.Contains(words[i], "(cAP") || strings.Contains(words[i], "(CaP"){
			if strings.Contains(words[i], "(cap,") || strings.Contains(words[i], "(Cap,") || strings.Contains(words[i], "(CAP,") || strings.Contains(words[i], "(cAP,") || strings.Contains(words[i], "(caP,") || strings.Contains(words[i], "(cAP,") || strings.Contains(words[i], "(CaP"){
				// Now initialize a variable to store the number that will be converted from a string to an integer
				num, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				if err != nil {
					fmt.Println("Error in Conversion", err)
				}
				// A contition that will only run if the number is less than or equal to the words from the begining upto i but not including i
				if num <= len(words[:i]) {
				for j := i - num; j < i; j++ {
					// only capitalize words at j the first letter while the rest remain lower case
					words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(string(words[j][1:]))
				}
				// Now append the words eliminating the i
				words = append(words[:i], words[i+2:]...)
				// another condition that checks if the number is greater than the words from the begining upto i but not including i
			} else if num > len(words[:i]) {
				for r := 0; r < len(words[:i]); r++ {
					words[r] = strings.ToUpper(string(words[r][0])) + strings.ToLower(string(words[r][1:]))
				}
				words = append(words[:i], words[i+2:]...)
			}
			// now handle the instance of "cap"
			} else {
				words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(string(words[i-1][1:]))
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	// return the words joined together with a space
	return strings.Join(words, " ")
}
