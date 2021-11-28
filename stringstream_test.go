package stream_test

import (
	"stream"
	"testing"
	"unicode/utf8"
)

func TestStringStream_Filter(t *testing.T) {
	tests := []struct {
		fn       func(string) bool
		data     []string
		expected []string
	}{
		{
			fn: func(s string) bool {
				return utf8.RuneCountInString(s) > 1
			},
			data:     []string{"1", "234", "哈哈哈", "黑", "XDD", "thx"},
			expected: []string{"234", "哈哈哈", "XDD", "thx"},
		},
		{
			fn: func(s string) bool {
				return utf8.RuneCountInString(s) == 1
			},
			data:     []string{"1", "234", "哈哈哈", "黑", "XDD", "thx"},
			expected: []string{"1", "黑"},
		},
	}

	for _, tt := range tests {
		actual := stream.NewStringStream(tt.data).
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

func TestStringStream_MapToInt(t *testing.T) {
	tests := []struct {
		fn       func(string) int
		data     []string
		expected []int
	}{
		{
			fn: func(s string) int {
				return len(s)
			},
			data:     []string{"1", "22", "555", "66666"},
			expected: []int{1, 2, 3, 5},
		},
		{
			fn: func(s string) int {
				return 0
			},
			data:     []string{"1", "22", "555", "66666"},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		actual := stream.NewStringStream(tt.data).
			MapToInt(tt.fn).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("MapToInt(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}

func TestStringStream_Collect(t *testing.T) {
	tests := []struct {
		data     []string
		expected []string
	}{
		{
			data:     []string{"1", "2", "3", "4", "5"},
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			data:     []string{"-1", "3", "12", "0"},
			expected: []string{"-1", "3", "12", "0"},
		},
		{
			data:     []string{"1", "1", "1"},
			expected: []string{"1", "1", "1"},
		},
	}

	for _, tt := range tests {
		actual := stream.NewStringStream(tt.data).
			Collect()
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Errorf("Collect(%v): expected %v, actual %v", tt.data, tt.expected, actual)
				return
			}
		}
	}
}
