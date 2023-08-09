package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestSelectFirst(t *testing.T) {
	type person struct {
		first string
		last  string
	}
	people := []person{
		{"hello", "world"},
		{"some", "person"},
		{"other", "name"},
	}

	peopleIter := iter.FromSlice(people)
	firstNameIter := iter.Select(peopleIter, func(p person) string { return p.first })
	first := iter.First(firstNameIter).UnwrapOr("")
	// equivalent
	/*
		first = ""
		for _, p := range people{
			first = p.first
			break
		}
	*/
	if first != "hello" {
		t.Fatalf("expected 'hello' found '%s'", first)
	}
}
