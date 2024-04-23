package mine

import (
	"reflect"
	"testing"
)

func TestApostrophe(t *testing.T) {
	sample := "' Alice '"
	expected := "'Alice'"

	result := Apostrophe(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
