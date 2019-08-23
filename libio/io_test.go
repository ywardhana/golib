package libio_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/golib/libio"
)

func TestReadLine(t *testing.T) {
	stringtest := "1 1"
	reader := bufio.NewReader(strings.NewReader(stringtest + "\n"))
	res := libio.ReadLine(reader)

	assert.Equal(t, stringtest, res)
}
