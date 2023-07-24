/*
This package provides generic iterators and iterator transformations in go.

Iterators defer iteration execution by pushing iteration logic into the Next() function.

This package has a dependency on [go-types](https://github.com/patrickhuber/go-types) for the Option[T any] type. Users of the iter module can utilize Result[T] to capture errors or perform transformations on slices that can contain errors.
*/
package iter
