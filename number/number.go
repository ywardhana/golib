package number

import (
	"reflect"
	"strconv"
)

const (
	DEFAULT_FLOAT_PREC = -1
	FLOAT_BITSIZE_32   = 32
	FLOAT_BITSIZE_64   = 64
	DEFAULT_FLOAT_FMT  = 'f'
	DEFAULT_BASE       = 10
	HEX_BASE           = 16
	BIN_BASE           = 2
)

// converting numbers to string
func ToString(param interface{}) (retVal string) {
	switch reflect.TypeOf(param).String() {
	case "int":
		retVal = strconv.FormatInt(int64(param.(int)), DEFAULT_BASE)
	case "int64":
		retVal = strconv.FormatInt(param.(int64), DEFAULT_BASE)
	case "uint":
		retVal = strconv.FormatUint(uint64(param.(uint)), DEFAULT_BASE)
	case "uint8":
		retVal = strconv.FormatUint(uint64(param.(uint8)), DEFAULT_BASE)
	case "uint16":
		retVal = strconv.FormatUint(uint64(param.(uint16)), DEFAULT_BASE)
	case "uint32":
		retVal = strconv.FormatUint(uint64(param.(uint32)), DEFAULT_BASE)
	case "uint64":
		retVal = strconv.FormatUint(param.(uint64), DEFAULT_BASE)
	case "float32":
		retVal = strconv.FormatFloat(float64(param.(float32)), DEFAULT_FLOAT_FMT, DEFAULT_FLOAT_PREC, FLOAT_BITSIZE_32)
	case "float64":
		retVal = strconv.FormatFloat(param.(float64), DEFAULT_FLOAT_FMT, DEFAULT_FLOAT_PREC, FLOAT_BITSIZE_64)
	}

	return
}

// converting non floating numbers to their hexadecimal representation in string
func ToHexString(param interface{}) (retVal string) {
	switch reflect.TypeOf(param).String() {
	case "int":
		retVal = strconv.FormatInt(int64(param.(int)), HEX_BASE)
	case "int64":
		retVal = strconv.FormatInt(param.(int64), HEX_BASE)
	case "uint":
		retVal = strconv.FormatUint(uint64(param.(uint)), HEX_BASE)
	case "uint8":
		retVal = strconv.FormatUint(uint64(param.(uint8)), HEX_BASE)
	case "uint16":
		retVal = strconv.FormatUint(uint64(param.(uint16)), HEX_BASE)
	case "uint32":
		retVal = strconv.FormatUint(uint64(param.(uint32)), HEX_BASE)
	case "uint64":
		retVal = strconv.FormatUint(param.(uint64), HEX_BASE)
	}
	return
}

// converting non floating numbers to their binary representation in string
func ToBinaryString(param interface{}) (retVal string) {
	switch reflect.TypeOf(param).String() {
	case "int":
		retVal = strconv.FormatInt(int64(param.(int)), BIN_BASE)
	case "int64":
		retVal = strconv.FormatInt(param.(int64), BIN_BASE)
	case "uint":
		retVal = strconv.FormatUint(uint64(param.(uint)), BIN_BASE)
	case "uint8":
		retVal = strconv.FormatUint(uint64(param.(uint8)), BIN_BASE)
	case "uint16":
		retVal = strconv.FormatUint(uint64(param.(uint16)), BIN_BASE)
	case "uint32":
		retVal = strconv.FormatUint(uint64(param.(uint32)), BIN_BASE)
	case "uint64":
		retVal = strconv.FormatUint(param.(uint64), BIN_BASE)
	}
	return
}
