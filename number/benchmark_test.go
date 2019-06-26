package number_test

import (
	"fmt"
	"testing"

	"github.com/ywardhana/golib/number"
)

var smallInt = 35
var bigInt = 999999999999999
var smallFloat = 3.5
var bigFloat = 999999999999999.999999999999999

func BenchmarkToStringSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", smallInt)
		_ = val
	}
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := number.ToString(smallInt)
		_ = val
	}
}

func BenchmarkToStringSprintfBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", bigInt)
		_ = val
	}
}

func BenchmarkToStringBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := number.ToString(bigInt)
		_ = val
	}
}

func BenchmarkToStringSprintfFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%f", bigFloat)
		_ = val
	}
}

func BenchmarkToStringBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := number.ToString(bigFloat)
		_ = val
	}
}
