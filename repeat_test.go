package iter_test

import (
	"reflect"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestRepeat(t *testing.T) {
	count := 8
	element := 10
	expected := make([]int, 0, count)
	for i := 0; i < count; i++ {
		expected = append(expected, element)
	}
	repeat := iter.Repeat(element, count)
	actual := iter.ToSlice(repeat)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("found %v but expected %v", actual, expected)
	}
}
