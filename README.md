Chain stream operations:

```go
r := stream.NewStringStream([]string{
        "", "1", "1", "333", "4444", "22", "333", "55555",
    }).
    MapToInt(func(s string) int {
        return len(s)
    }).
    Distinct().
    Sorted().
    Peek(func(i int) {
        fmt.Printf("%v ", i)
    }).
    Filter(func(i int) bool {
        return i > 0
    }).
    Skip(1).
    Limit(5).
    Reduce(0, func(i int, i2 int) int {
        return i + i2
    })
fmt.Printf("\nresult: %v", r)

// Output:
// 0 1 2 3 4 5
// result: 14
```
