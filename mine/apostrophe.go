package mine

import "strings"

func Apostrophe(str string) string {
	str = strings.ReplaceAll(str, "' ", " '")
	str = strings.ReplaceAll(str, " '", "'")

	// Trim any leading or trailing whitespace to maintain clean output.
	return strings.TrimSpace(str)
}
