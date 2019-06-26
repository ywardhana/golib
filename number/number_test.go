package number_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/number"
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
		assert.Equal(t, tt.Expected, number.ToString(tt.Value), tt.Title)
	}

}

func TestToHexString(t *testing.T) {

	tests := []struct {
		Title    string
		Value    interface{}
		Expected string
	}{
		{
			Title:    "Test int",
			Value:    123,
			Expected: fmt.Sprintf("%x", 123),
		},
		{
			Title:    "Test int64",
			Value:    int64(999999999999999999),
			Expected: fmt.Sprintf("%x", 999999999999999999),
		},
		{
			Title:    "Test uint",
			Value:    uint(123),
			Expected: fmt.Sprintf("%x", 123),
		},
		{
			Title:    "Test uint8",
			Value:    uint8(123),
			Expected: fmt.Sprintf("%x", 123),
		},
		{
			Title:    "Test uint16",
			Value:    uint16(123),
			Expected: fmt.Sprintf("%x", 123),
		},
		{
			Title:    "Test uint32",
			Value:    uint32(123),
			Expected: fmt.Sprintf("%x", 123),
		},
		{
			Title:    "Test uint64",
			Value:    uint64(999999999999999999),
			Expected: fmt.Sprintf("%x", 999999999999999999),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, number.ToHexString(tt.Value), tt.Title)
	}

}

func TestToBinaryString(t *testing.T) {

	tests := []struct {
		Title    string
		Value    interface{}
		Expected string
	}{
		{
			Title:    "Test int",
			Value:    123,
			Expected: fmt.Sprintf("%b", 123),
		},
		{
			Title:    "Test int64",
			Value:    int64(999999999999999999),
			Expected: fmt.Sprintf("%b", 999999999999999999),
		},
		{
			Title:    "Test uint",
			Value:    uint(123),
			Expected: fmt.Sprintf("%b", 123),
		},
		{
			Title:    "Test uint8",
			Value:    uint8(123),
			Expected: fmt.Sprintf("%b", 123),
		},
		{
			Title:    "Test uint16",
			Value:    uint16(123),
			Expected: fmt.Sprintf("%b", 123),
		},
		{
			Title:    "Test uint32",
			Value:    uint32(123),
			Expected: fmt.Sprintf("%b", 123),
		},
		{
			Title:    "Test uint64",
			Value:    uint64(999999999999999999),
			Expected: fmt.Sprintf("%b", 999999999999999999),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, number.ToBinaryString(tt.Value), tt.Title)
	}

}
