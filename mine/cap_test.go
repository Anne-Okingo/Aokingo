package mine

import (
	"reflect"
	"testing"
)

func TestCap(t *testing.T) {
	sample := "Alice has two (cap) Kenyan Shilling (cap, 2) in her bag."
	expected := "Alice has Two Kenyan Shilling in her bag."

	result := Cap(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Cap failed. Expected: %v. Got: %v", expected, result)
	}
}
