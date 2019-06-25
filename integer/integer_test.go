package integer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/integer"
)

func TestToString(t *testing.T) {

	tests := []struct {
		Title    string
		Value    interface{}
		Expected string
	}{
		{
			Title:    "Test int",
			Value:    123,
			Expected: "123",
		},
		{
			Title:    "Test int64",
			Value:    int64(999999999999999999),
			Expected: "999999999999999999",
		},
		{
			Title:    "Test uint",
			Value:    uint(123),
			Expected: "123",
		},
		{
			Title:    "Test uint8",
			Value:    uint8(123),
			Expected: "123",
		},
		{
			Title:    "Test uint16",
			Value:    uint16(123),
			Expected: "123",
		},
		{
			Title:    "Test uint32",
			Value:    uint32(123),
			Expected: "123",
		},
		{
			Title:    "Test uint64",
			Value:    uint64(999999999999999999),
			Expected: "999999999999999999",
		},
		{
			Title:    "Test float32",
			Value:    float32(1.2),
			Expected: "1.2",
		},
		{
			Title:    "Test float64",
			Value:    1.2,
			Expected: "1.2",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, integer.ToString(tt.Value), tt.Title)
	}

}
