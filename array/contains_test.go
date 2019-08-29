package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ywardhana/golib/slice"
)

func TestContains(t *testing.T) {
	tests := []struct {
		array    interface{}
		item     interface{}
		expected bool
	}{
		{
			array:    []int{1, 2, 3, 4, 5},
			item:     5,
			expected: true,
		},
		{
			array:    []int{1, 2, 3, 4},
			item:     5,
			expected: false,
		},
		{
			array:    []interface{}{1, "haha", 3, 4.5},
			item:     5,
			expected: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, slice.Contains(tt.array, tt.item))
	}
}
