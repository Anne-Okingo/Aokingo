package mine

import (
	"reflect"
	"testing"
)

func TestLow(t *testing.T) {
	sample := "ALICE HAS A BRIGHT (low, 2) FUTURE (low)"
	expected := "ALICE HAS a bright future"

	result := Low(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Low Failed Expected : %v. Got : %v .", expected, result)
	}
}
