package iter_test

import (
	"testing"
	"time"

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

func TestTransforms(t *testing.T) {
	type address struct{}
	type phone struct{}
	type person struct {
		FirstName   string
		LastName    string
		DateOfBirth time.Time
		Addresses   []address
		Phones      []phone
	}

	people := iter.ToSlice(iter.New(person{
		Addresses: iter.ToSlice(iter.New(
			address{},
			address{})),
		Phones: iter.ToSlice(iter.New(
			phone{},
			phone{})),
	}))
	if len(people) != 1 {
		t.Fatalf("expected length to be 1, found %d", len(people))
	}
}
