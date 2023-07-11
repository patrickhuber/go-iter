# Iteration library in go

This package provides generic iterators and iterator transformations in go.

Iterators defer iteration execution by pushing iteration logic into the Next() function.

This package has a dependency on "github.com/patrickhuber/go-types" for the Option[T any] type. Users of the iter module can utilize Result[T] to capture errors or perform transformations on slices that can contain errors. 

# Range

The [Range](range.go) function creates an iterator over `[begin..end)` (inclusive begin, exclusive end)

```golang
rng := iter.Range(0, 10)
iter.ForEachIndex(rng, func(index int, i int) {
    if index > 0 {
        fmt.Print(" ")
    }
    fmt.Print(i)
})
// prints:
// 0 1 2 3 4 5 6 7 8 9
```

# Select

The [Select](select.go) function transforms an iterator of one type to an iterator of another type. 

```golang
rng := iter.Range(0,10)
strRng := iter.Select(rng, strconv.Itoa)
iter.ForEachIndex(strRng, func(index int, s string){
    if index > 0{
        fmt.Print(" ")
    }
    fmt.Print("'%s'", s)
})
// prints : '0' '1' '2' '3' '4' '5' '6' '7' '8' '9'
```

# Where

The [Where](where.go) function removes items from an iterator using a predicate function.

```golang
rng := iter.Range(0,10)
even := func(i int){ return i % 2 }
evens := iter.Where(rng, even)
iter.ForEachIndex(evens, func(index int, i int){
    if index > 0{
        fmt.Prin(" ")
    }
    fmt.Print("%d", i)
})
// prints : 0 2 4 6 8
```

# FromSlice

The [FromSlice](slice.go) function returns a slice iterator over the given slice

```golang
slice := []int{1, 3, 5, 7, 9}
it := iter.FromSlice(slice)
iter.ForEachIndex(it, func(index int, i int){
    if index > 0{
        fmt.Prin(" ")
    }
    fmt.Print("%d", i)
})
// prints 1 3 5 7 9
```

# ToSlice

The [ToSlice](slice.go) function returns a slice from the given iterator by iterating over all elements.

```golang
rng := iter.Range(0, 10)
slice := iter.ToSlice(rng)
fmt.Println(slice)
// prints : [0 1 2 3 4 5 6 7 8 9]
```

# FromMap

The [FromMap](map.go) function returns an iterator over the given map. FromMap captures keys of the map in the first call to `Next()`. Subsequent calls to `Next()` track an index into the key slice to return the next key value pair. Each call to `Next()` returns an Option tuple of Key Value pairs.

```golang
m := map[string]int{"0":0, "1":1, "2":2, "3":3}
it := iter.FromMap(m)
fmt.Print("[ ")
iter.ForEachIndex(it, func(index int, tup types.Tuple2[string, int]){
    if index > 0{
        fmt.Print(", ")
    }    
    k, v := tup.Deconstruct()
    fmt.Print("'%s':%d", k, v)
})
fmt.Print(" ]")
// prints : [ '0':0, '1':1, '2':2, '3':3 ]
```
