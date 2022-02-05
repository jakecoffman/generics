# generics

JavaScript style generics package.

For usage see [main.go](main.go)

## Performance

Everything in this library is slower than if written by hand, but the readability might be worth it to you in your use case.

## Design choices

### index vs pointers

Where JavaScript does this:

```javascript
const things = [{id: 1},{id: 2}]
things.find(e => e.id === 2)
```

We would do this in Go, passing a pointer to Thing to avoid a potentially expensive copy:

```go
things := []Thing{{id: 1}, {id: 2}}
slices.Find(things, func(e *Thing) bool { return e.id == 2 })
```

However, I think this may be confusing to JS developers if the array was a primitive:

```go
arr := []int{1, 2}
slices.Find(arr, func(e *int) bool { return *e == 2 })
```

Therefore, in a lot of places the function receives the index to the array rather than a pointer.

## Chainability

Chaining operations doesn't work well in Go unless we introduce a higher level concept like a Stream. I think we've all seen the Java Streams and that's just not as easy to use as a simple function.

You can avoid making a ton of variables by reusing where possible:

```go
things := []int{1, 2, 3, 4}
things = slices.Filter(things, func(i int) bool { return things[i]%2 == 0 })
things = slices.Map(things, func(i int) int { return i + 1 })
```

Obviously if Map changes the type this won't work.
