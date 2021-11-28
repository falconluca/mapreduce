# Stream

Go实现Java Stream风格的函数式编程API

## Usage

链式调用

```go
_ = stream.NewIntStream([]int{5, 6, -1, 1, 1, 1, 1, 1, 4, 2, 3}).
	Distinct().
	Sorted().
	Skip(1).
	Limit(5).
	Peek(func(i int) {
	    fmt.Printf("%v ", i)
	}).
	Collect()
```

函数式编程三架马车: Map, Filter, Reduce

```go
r := stream.NewStringStream([]string{
	    "", "1", "22", "333", "4444", "55555",
	}).
	MapToInt(func(s string) int {
	    return len(s)
	}).
	Filter(func(i int) bool {
	    return i > 0
	}).
	Reduce(0, func(i int, i2 int) int {
	    return i + i2
	})
	fmt.Println(r)

// Output:
// 15
```
