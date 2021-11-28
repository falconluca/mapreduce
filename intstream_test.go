package stream_test

import (
	"fmt"
	"stream"
	"testing"
)

func TestIntStream_MapToInt(t *testing.T) {
	tests := []struct {
		fn       func(int) int
		data     []int
		expected []int
	}{
		{
			fn: func(i int) int {
				return i * 2
			},
			data:     []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			fn: func(i int) int {
				return i + 2
			},
			data:     []int{1, 2, 3, 4, 5},
			expected: []int{3, 4, 5, 6, 7},
		},
		{
			fn: func(_ int) int {
				return 0
			},
			data:     []int{4, 2, 9, 2, -13},
			expected: []int{0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			MapToInt(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("MapToInt(%d): expected %d, actual %d", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_MapToString(t *testing.T) {
	tests := []struct {
		fn       func(int) string
		data     []int
		expected []string
	}{
		{
			fn: func(i int) string {
				return fmt.Sprintf("%v", i)
			},
			data:     []int{1, 2, 5, 6, 7},
			expected: []string{"1", "2", "5", "6", "7"},
		},
		{
			fn: func(i int) string {
				return fmt.Sprintf("Greetings! %v", i)
			},
			data:     []int{1, 2, 5, 6, 7},
			expected: []string{"Greetings! 1", "Greetings! 2", "Greetings! 5", "Greetings! 6", "Greetings! 7"},
		},
		{
			fn: func(i int) string {
				return ""
			},
			data:     []int{1, 2, 5, 6, 7},
			expected: []string{"", "", "", "", ""},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			MapToString(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("MapToString(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Filter(t *testing.T) {
	tests := []struct {
		fn       func(int) bool
		data     []int
		expected []int
	}{
		{
			fn: func(i int) bool {
				return i > 100
			},
			data:     []int{-12, 2, 3, 100, 120, 123, 430},
			expected: []int{120, 123, 430},
		},
		{
			fn: func(i int) bool {
				return i > -12 && i != 0
			},
			data:     []int{-12, -2, 0, 3, 100, 120, 123, 430},
			expected: []int{-2, 3, 100, 120, 123, 430},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Filter(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Filter(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Peek(t *testing.T) {
	tests := []struct {
		fn       func(int)
		data     []int
		expected []int
	}{
		{
			fn: func(i int) {
				fmt.Printf("Greetings! %v", i)
			},
			data:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			fn: func(i int) {
				fmt.Printf("==> %v", i)
			},
			data:     []int{-1, 4, 5},
			expected: []int{-1, 4, 5},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Peek(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Peek(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_ForEach(t *testing.T) {
	// TODO
}

func TestIntStream_Limit(t *testing.T) {
	type ExpectedType struct {
		data []int
		size int
	}
	tests := []struct {
		limit    int
		data     []int
		expected ExpectedType
	}{
		{
			limit: 1,
			data:  []int{1, 2, 4, 5},
			expected: ExpectedType{
				data: []int{1},
				size: 1,
			},
		},
		{
			limit: 5,
			data:  []int{1, 2, 4, 5, 6},
			expected: ExpectedType{
				data: []int{1, 2, 4, 5, 6},
				size: 5,
			},
		},
		{
			limit: 10,
			data:  []int{1, 2, 4, 5, 6},
			expected: ExpectedType{
				data: []int{1, 2, 4, 5, 6},
				size: 5,
			},
		},
		{
			limit: 10,
			data:  []int{1, 2, 4, 5, 6, 1, 2, 4, 5, 6, 1, 2, 4, 5, 6},
			expected: ExpectedType{
				data: []int{1, 2, 4, 5, 6, 1, 2, 4, 5, 6},
				size: 10,
			},
		},
		{
			limit: 0,
			data:  []int{1, 2, 4, 5},
			expected: ExpectedType{
				data: []int{},
				size: 0,
			},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Limit(tt.limit).
			Collect()
		if len(actual) != tt.expected.size {
			t.Errorf("Limit(%v): slice size expected %v, actual %v", tt.data, tt.limit, len(actual))
			return
		}
		for i := range tt.expected.data {
			if actual[i] != tt.expected.data[i] {
				t.Errorf("Limit(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Skip(t *testing.T) {
	type ExpectedType struct {
		data []int
		size int
	}
	tests := []struct {
		cnt      int
		data     []int
		expected ExpectedType
	}{
		{
			cnt:  1,
			data: []int{1, 2, 3, 5},
			expected: ExpectedType{
				data: []int{2, 3, 5},
				size: 3,
			},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Skip(tt.cnt).
			Collect()
		if len(actual) != tt.expected.size {
			t.Errorf("Skip(%v): slice size expected %v, actual %v", tt.data, tt.expected.size, len(actual))
			return
		}
		for i := range tt.expected.data {
			if actual[i] != tt.expected.data[i] {
				t.Errorf("Skip(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Distinct(t *testing.T) {
	tests := []struct {
		data     []int
		expected []int
	}{
		{
			data:     []int{1, 2, 2, 2, 3, 4, 5, 6, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			data:     []int{1, 1, 1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Distinct().
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Distinct(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Sorted(t *testing.T) {
	tests := []struct {
		data     []int
		expected []int
	}{
		{
			data:     []int{4, 2, 5, 1},
			expected: []int{1, 2, 4, 5},
		},
		{
			data:     []int{4, -1, 5, 1},
			expected: []int{-1, 1, 4, 5},
		},
		{
			data:     []int{4, 2, 5, 1, 9, 0, -13},
			expected: []int{-13, 0, 1, 2, 4, 5, 9},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Sorted().
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Sorted(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_SortedBy(t *testing.T) {
	tests := []struct {
		data     []int
		expected []int
		fn       func(int, int) bool
	}{
		{
			data:     []int{12, 2, 34, 0, -1},
			expected: []int{-1, 0, 2, 12, 34},
			fn: func(i int, i2 int) bool {
				return i < i2
			},
		},
		{
			data:     []int{12, 2, 34, 0, -1},
			expected: []int{34, 12, 2, 0, -1},
			fn: func(i int, i2 int) bool {
				return i > i2
			},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			SortedBy(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("SortedBy(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Collect(t *testing.T) {
	tests := []struct {
		data     []int
		expected []int
	}{
		{
			data:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			data:     []int{-1, 3, -12, 0},
			expected: []int{-1, 3, -12, 0},
		},
		{
			data:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Collect(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestIntStream_Count(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			data:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			data:     []int{-1, 3, -12, 0},
			expected: 4,
		},
		{
			data:     []int{1},
			expected: 1,
		},
		{
			data:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Count()
		if actual != tt.expected {
			t.Errorf("Count(%v): expected %v, actual %v", tt.data, tt.expected, actual)
			return
		}
	}
}

func TestIntStream_Max(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			data:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			data:     []int{-1, 3, -12, 0},
			expected: 3,
		},
		{
			data:     []int{1},
			expected: 1,
		},
		{
			data:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Max()
		if actual != tt.expected {
			t.Errorf("Max(%v): expected %v, actual %v", tt.data, tt.expected, actual)
			return
		}
	}
}

func TestIntStream_Min(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			data:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			data:     []int{-1, 3, -12, 0},
			expected: -12,
		},
		{
			data:     []int{1},
			expected: 1,
		},
		{
			data:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Min()
		if actual != tt.expected {
			t.Errorf("Min(%v): expected %v, actual %v", tt.data, tt.expected, actual)
			return
		}
	}
}

func TestIntStream_Reduce(t *testing.T) {
	tests := []struct {
		initial  int
		fn       func(int, int) int
		data     []int
		expected int
	}{
		{
			initial: 0,
			fn: func(i int, i2 int) int {
				return i + i2
			},
			data:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			initial: 1,
			fn: func(i int, i2 int) int {
				return i * i2
			},
			data:     []int{2, 3, 4, 5},
			expected: 120,
		},
	}

	for _, tt := range tests {
		actual := stream.NewIntStream(tt.data).
			Reduce(tt.initial, tt.fn)
		if actual != tt.expected {
			t.Errorf("Reduce(%v): expected %v, actual %v", tt.data, tt.expected, actual)
			return
		}
	}
}
