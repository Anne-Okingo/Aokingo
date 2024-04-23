package mine

import (
	"reflect"
	"testing"
)

func TestHex(t *testing.T) {
	sample := "Alice has 1E (hex) Kenyan Shilling"
	expected := "Alice has 30 Kenyan Shilling"

	result := Hex(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Hex failed. Expected: %v. Got: %v", expected, result)
	}
}
