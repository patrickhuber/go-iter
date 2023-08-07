package iter_test

import (
	"reflect"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestRunes(t *testing.T) {
	str := "hello world"
	expected := []rune(str)
	rs := iter.Runes(str)
	actual := iter.ToSlice(rs)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%s' to equal '%s'", string(actual), string(expected))
	}
}

func TestUnicodeRunes(t *testing.T) {
	str := "สวัสดี" // hello in thai
	expected := []rune(str)
	rs := iter.Runes(str)
	actual := iter.ToSlice(rs)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%s' to equal '%s'", string(actual), string(expected))
	}
}
