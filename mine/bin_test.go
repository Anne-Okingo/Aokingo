package mine

import (
	"reflect"
	"testing"
)

func TestBin(t *testing.T) {
	sample := "Alice has 10 (bin) Kenyan Shilling"
	expected := "Alice has 2 Kenyan Shilling"

	result := Bin(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Bin failed. Expected: %v. Got: %v", expected, result)
	}
}
