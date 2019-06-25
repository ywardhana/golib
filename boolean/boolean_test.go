package boolean_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/boolean"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		Title    string
		Bool     bool
		Expected int
	}{
		{
			Title:    "false",
			Bool:     false,
			Expected: 0,
		},
		{
			Title:    "true",
			Bool:     true,
			Expected: 1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, boolean.ToInt(tt.Bool), tt.Title)
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		Title    string
		Bool     bool
		Expected string
	}{
		{
			Title:    "false",
			Bool:     false,
			Expected: "false",
		},
		{
			Title:    "true",
			Bool:     true,
			Expected: "true",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, boolean.ToString(tt.Bool), tt.Title)
	}
}
