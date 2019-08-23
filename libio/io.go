package libio

import (
	"bufio"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
