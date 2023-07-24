# Iteration library in go

This package provides generic iterators and iterator transformations in go.

Iterators defer iteration execution by pushing iteration logic into the Next() function.

This package has a dependency on [go-types](https://github.com/patrickhuber/go-types) for the Option[T any] type. Users of the iter module can utilize Result[T] to capture errors or perform transformations on slices that can contain errors. 

# Looping

Here are three ways to loop over an iterator

## [For Loop](for_test.go)

A for loop has init, condition and post sections that can be used on an iterator.

```golang
rng := iter.Range(0, 10)
for op := rng.Next(); op.IsSome(); op = rng.Next() {
    // code in here
}
```

## [Async](async.go) (calling range on a channel)

You can also loop over a channel created from an Iterator[T] using the iter.Async function. You can specify a context with the iter.WithContext[T] option. If no context is specified, context.Background() is used.

```golang
rng := iter.Range(0, 10)
for value := range iter.Async(rng) {
    // code in here
}
```

You can also get an index 

```golang
rng := iter.Range(0, 10)
for i, value := range iter.Async(rng){
    // code in here
}
```

Specify a context

```golang
rng := iter.Range(0,10)
op := iter.WithContext[int](context.TODO())
for value := range iter.Async(rng, op){
    // code here
}
```

## [ForEach](for_each.go)

ForEach uses an anonymous function to iterate over each item to the caller.

```golang
rng := iter.Range(0, 10)
iter.ForEach(rng, func(i int) {
    fmt.Print(" ")
    fmt.Print(i)
})
// prints:
//  0 1 2 3 4 5 6 7 8 9
```

## [ForEachIndex](for_each.go)

ForEachIndex uses an anonymous function to iterate over each item and the index of that item to the caller.

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

# [Range](range.go)

The Range function creates an iterator over `[begin..end)` (inclusive begin, exclusive end)

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

# [Repeat](repeat.go)

Repeat returns an iterator that repeats the given `element`, `count` times

```golang
rep := iter.Repeat(10, 5)
iter.ForEachIndex(rng, func(index int, i int) {
    if index > 0 {
        fmt.Print(" ")
    }
    fmt.Print(i)
})
// prints:
// 10 10 10 10 10
```

# [Select](select.go)

The Select function transforms an iterator of one type to an iterator of another type. 

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

# [Where](where.go)

The Where function removes items from an iterator using a predicate function.

```golang
rng := iter.Range(0,10)
even := func(i int){ return i % 2 }
evens := iter.Where(rng, even)
iter.ForEachIndex(evens, func(index int, i int){
    if index > 0{
        fmt.Print(" ")
    }
    fmt.Print("%d", i)
})
// prints : 0 2 4 6 8
```

# [FromSlice](slice.go)

The FromSlice function returns a slice iterator over the given slice

```golang
slice := []int{1, 3, 5, 7, 9}
it := iter.FromSlice(slice)
iter.ForEachIndex(it, func(index int, i int){
    if index > 0{
        fmt.Print(" ")
    }
    fmt.Print("%d", i)
})
// prints 1 3 5 7 9
```

# [ToSlice](slice.go) 

The ToSlice function returns a slice from the given iterator by iterating over all elements.

```golang
rng := iter.Range(0, 10)
slice := iter.ToSlice(rng)
fmt.Println(slice)
// prints : [0 1 2 3 4 5 6 7 8 9]
```

# [FromMap](map.go) 

The FromMap function returns an iterator over the given map. FromMap captures keys of the map in the first call to `Next()`. Subsequent calls to `Next()` track an index into the key slice to return the next key value pair. Each call to `Next()` returns an Option tuple of Key Value pairs.

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

# [FromMapAsync](map.go)

The FromMapAsync function returns an iterator over the given map using channels and the supplied context. FromMapAsync uses a go routine to iterate over the map and channels to capture the key value pair. The `Next()` function removes a tuple key, value from the channel.

```golang
m := map[string]int{"0":0, "1":1, "2":2, "3":3}
it := iter.FromMapAsync(m, context.Background())
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


# [Count](count.go) 

The Count function returns the count of element in the iterator by iterating over all the elements. 

```golang
expected := 10
rng := iter.Range(0, expected)
actual := iter.Count(rng)
if actual != expected {
    t.Fatalf("expected count of %d but found %d", expected, actual)
}
```

# [FromChannel](channel.go)

The FromChannel function returns an iterator over the given channel. A context can be specified in options for early termination.

```golang
ch := make(chan int)
go func(c chan int) {
    defer close(c)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}(ch)
it := iter.FromChannel(ch)
iter.ForEach(it, func(i int) {
    fmt.Println(i)
})
```