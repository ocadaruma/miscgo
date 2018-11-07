package miscgo

import (
	"sort"
)

type SuffixArray struct {
	indices []int
}

type suffix struct {
	s string
	index int
}
func NewSufficArray(s string) *SuffixArray {
	nullEnded := s + "\x00"
	n := len(nullEnded)

	buffer := make([]*suffix, n)
	for i := 0; i < n; i++ {
		buffer[i] = &suffix{
			s: nullEnded[i:n],
			index: i,
		}
	}

	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].s < buffer[j].s
	})

	indices := make([]int, n)
	for i, suffix := range buffer {
		indices[suffix.index] = i
	}

	return &SuffixArray{
		indices: indices,
	}
}

func (self *SuffixArray) At(i int) int {
	return self.indices[i]
}
