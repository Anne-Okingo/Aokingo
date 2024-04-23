package mine

import (
	"testing"
	"reflect"
)

func TestUp (t *testing.T) {
	sample := "Alice has two (up) kenyan shilling (up, 2) in her bag."
	expected := "Alice has TWO KENYAN SHILLING in her bag."

	result := Up(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Cap failed. Expected: %v. Got: %v", expected, result)
	}
}