package mine

import (
	"reflect"
	"testing"
)

func TestLow(t *testing.T) {
	sample := "Alice has TWO (low) KENYAN SHILLING (low, 2) in her bag."
	expected := "Alice has two kenyan shilling in her bag."

	result := Low(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Low failed. Expected: %v. Got: %v", expected, result)
	}
}
