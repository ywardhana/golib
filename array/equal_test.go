package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/array"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		arrA     interface{}
		arrB     interface{}
		expected bool
	}{
		{
			arrA:     []int{1, 2, 3, 4},
			arrB:     []int{1, 2, 3, 4},
			expected: true,
		},
		{
			arrA:     []int{4, 2, 1, 3},
			arrB:     []int{1, 2, 3, 4},
			expected: true,
		},
		{
			arrA:     []int{4, 2, 1},
			arrB:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			arrA:     []int{4, 2, 1, 3, 1},
			arrB:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			arrA:     []int{4, 2, 1, 3},
			arrB:     []int{1, 1, 2, 3, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, array.Equal(tt.arrA, tt.arrB))
	}
}
