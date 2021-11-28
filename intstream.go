package stream

import (
	"math"
	"sort"
)

type IntStream struct {
	data []int
}

func NewIntStream(data []int) *IntStream {
	return &IntStream{
		data,
	}
}

func (s *IntStream) MapToInt(fn func(int) int) *IntStream {
	var result []int
	for _, item := range s.data {
		result = append(result, fn(item))
	}
	return &IntStream{
		result,
	}
}

func (s *IntStream) MapToString(fn func(int) string) *StringStream {
	var result []string
	for _, item := range s.data {
		result = append(result, fn(item))
	}
	return NewStringStream(result)
}

func (s *IntStream) Filter(fn func(int) bool) *IntStream {
	var result []int
	for _, item := range s.data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return &IntStream{
		result,
	}
}

func (s *IntStream) Peek(fn func(int)) *IntStream {
	for _, item := range s.data {
		fn(item)
	}
	return s
}

func (s *IntStream) Limit(limit int) *IntStream {
	var result []int
	arr := s.data
	for i := 0; i < len(arr); i++ {
		if i > (limit - 1) {
			break
		}
		result = append(result, arr[i])
	}
	return &IntStream{
		result,
	}
}

func (s *IntStream) Skip(cnt int) *IntStream {
	var result []int
	arr := s.data
	for i := 0; i < len(arr); i++ {
		if i < cnt {
			continue
		}
		result = append(result, arr[i])
	}
	return &IntStream{
		result,
	}
}

func (s *IntStream) Distinct() *IntStream {
	keys := make(map[int]bool)
	var result []int
	for _, item := range s.data {
		if _, ok := keys[item]; ok {
			continue
		}
		keys[item] = true
		result = append(result, item)
	}
	return &IntStream{
		result,
	}
}

func (s *IntStream) Sorted() *IntStream {
	result := s.data
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return &IntStream{
		result,
	}
}

func (s *IntStream) SortedBy(fn func(int, int) bool) *IntStream {
	result := s.data
	sort.SliceStable(result, func(i, j int) bool {
		return fn(result[i], result[j])
	})
	return &IntStream{
		result,
	}
}

func (s *IntStream) Reduce(initial int, fn func(int, int) int) int {
	result := initial
	arr := s.data
	for i := 0; i < len(arr); i++ {
		result = fn(result, arr[i])
	}
	return result
}

func (s *IntStream) ForEach(fn func(int)) {
	for _, item := range s.data {
		fn(item)
	}
}

func (s *IntStream) Collect() []int {
	return s.data
}

func (s *IntStream) Count() int {
	return len(s.data)
}

func (s *IntStream) Max() int {
	var result int
	arr := s.data
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result = arr[i]
			continue
		}
		max := math.Max(float64(arr[i]), float64(result))
		result = int(max)
	}
	return result
}

func (s *IntStream) Min() int {
	var result int
	arr := s.data
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result = arr[i]
			continue
		}
		min := math.Min(float64(arr[i]), float64(result))
		result = int(min)
	}
	return result
}
