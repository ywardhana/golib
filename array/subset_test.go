package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/array"
)

func TestSubset(t *testing.T) {
	tests := []struct {
		array    interface{}
		subArray interface{}
		expected bool
	}{
		{
			array:    []int{1, 2, 3, 4, 5},
			subArray: []int{1, 2, 3},
			expected: true,
		},
		{
			array:    []int{1, 2, 3, 4},
			subArray: []int{5},
			expected: false,
		},
		{
			array:    []int{1, 2, 3, 4},
			subArray: []int{1, 1, 2, 3},
			expected: false,
		},
		{
			array:    []interface{}{1, "haha", 3, 4.5},
			subArray: []interface{}{5},
			expected: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, array.Subset(tt.subArray, tt.array))
	}
}
