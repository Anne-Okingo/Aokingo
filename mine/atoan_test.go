package mine

import (
	"reflect"
	"testing"
)

func TestAToAn(t *testing.T) {
    input := "I am a author"
    expected := "I am an author"

    result := AToAn(input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Test AToAn failed. Expected: %v. Got: %v", expected, result)
    }
}