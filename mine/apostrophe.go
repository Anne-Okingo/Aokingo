package mine

import (
	"strings"
)

func Apostrophe(str string) string {
	str = strings.ReplaceAll(str, "' ", " '")
	str = strings.ReplaceAll(str, " '", "'")

	return strings.TrimSpace(str)
}
