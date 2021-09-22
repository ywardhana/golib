package array_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/array"
)

func TestMap(t *testing.T) {
	tests := []struct {
		array    interface{}
		expected interface{}
		callback func(e interface{}) (interface{}, error)
		err      error
	}{
		{
			array:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 10},
			callback: func(e interface{}) (interface{}, error) {
				return e.(int) * 2, nil
			},
			err: nil,
		},
		{
			array:    []string{"palu", "pala", "pale", "polo"},
			expected: []string{"pau", "paa", "pae", "poo"},
			callback: func(e interface{}) (interface{}, error) {
				return strings.ReplaceAll(e.(string), "l", ""), nil
			},
			err: nil,
		},
		{
			array:    []string{"palu", "pala", "pale", "polo"},
			expected: nil,
			callback: func(e interface{}) (interface{}, error) {
				return nil, errors.New("there is an error")
			},
			err: errors.New("there is an error"),
		},
	}

	for _, tt := range tests {
		result, err := array.Map(tt.array, tt.callback)
		assert.Equal(t, tt.err, err)
		if tt.expected != nil {
			assert.True(t, array.Equal(tt.expected, result))
		} else {
			assert.Equal(t, tt.expected, result)
		}

	}
}
