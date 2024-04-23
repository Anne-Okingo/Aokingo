package mine

import (
	"reflect"
	"testing"
)

func TestPunc(t *testing.T) {
	sample := "I was sitting over there ,and then BAMM !! I was thinking ... You were right"
	expected := "I was sitting over there, and then BAMM!! I was thinking... You were right"

	result := Punc(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Cap failed. Expected: %v. Got: %v", expected, result)
	}
}
