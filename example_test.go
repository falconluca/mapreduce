package stream_test

import (
	"fmt"
	"stream"
)

func ExampleIntStream() {
	_ = stream.NewIntStream([]int{5, 6, -1, 1, 1, 1, 1, 1, 4, 2, 3}).
		Distinct().
		Sorted().
		Skip(1).
		Limit(5).Peek(func(i int) {
		fmt.Printf("%v ", i)
	}).
		Collect()

	// Output:
	// 1 2 3 4 5
}

func ExampleNotAffectEachOther() {
	mid := stream.NewIntStream([]int{1, 2, 3, 4, 5, 6}).
		SortedBy(func(i int, i2 int) bool {
			return i > i2
		})
	prev := mid.MapToInt(func(i int) int {
		return i * 2
	}).
		Collect()
	cur := mid.MapToInt(func(i int) int {
		return i * 2
	}).
		Collect()
	fmt.Println(prev)
	fmt.Println(cur)

	// Output:
	// [12 10 8 6 4 2]
	// [12 10 8 6 4 2]
}
