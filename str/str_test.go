package str_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/str"
)

func TestReverse(t *testing.T) {
	strTest := "test"
	reverseStr := "tset"
	result := str.Reverse(strTest)
	assert.Equal(t, reverseStr, result)
}

func TestContains(t *testing.T) {
	arr := []string{"satu", "dua", "tiga"}
	assert.True(t, str.Contains(arr, "satu"))
	assert.False(t, str.Contains(arr, "empat"))
}

func TestToSnakeCase(t *testing.T) {
	result := str.ToSnakeCase("satuDuaTiga")
	assert.Equal(t, result, "satu_dua_tiga")
}
