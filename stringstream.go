package stream

type StringStream struct {
	data []string
}

func NewStringStream(data []string) *StringStream {
	return &StringStream{
		data,
	}
}

func (s *StringStream) Filter(fn func(i string) bool) *StringStream {
	var result []string
	for _, item := range s.data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return NewStringStream(result)
}

func (s *StringStream) MapToInt(fn func(string2 string) int) *IntStream {
	var result []int
	for _, item := range s.data {
		result = append(result, fn(item))
	}
	return NewIntStream(result)
}

func (s *StringStream) Collect() []string {
	return s.data
}
