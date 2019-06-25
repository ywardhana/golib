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
)

func ToString(intVal interface{}) (retVal string) {
	switch reflect.TypeOf(intVal).String() {
	case "int":
		retVal = strconv.Itoa(intVal.(int))
	case "int64":
		retVal = strconv.FormatInt(intVal.(int64), DEFAULT_BASE)
	case "uint":
		retVal = strconv.FormatUint(uint64(intVal.(uint)), DEFAULT_BASE)
	case "uint8":
		retVal = strconv.FormatUint(uint64(intVal.(uint8)), DEFAULT_BASE)
	case "uint16":
		retVal = strconv.FormatUint(uint64(intVal.(uint16)), DEFAULT_BASE)
	case "uint32":
		retVal = strconv.FormatUint(uint64(intVal.(uint32)), DEFAULT_BASE)
	case "uint64":
		retVal = strconv.FormatUint(intVal.(uint64), DEFAULT_BASE)
	case "float32":
		retVal = strconv.FormatFloat(float64(intVal.(float32)), DEFAULT_FLOAT_FMT, DEFAULT_FLOAT_PREC, FLOAT_BITSIZE_32)
	case "float64":
		retVal = strconv.FormatFloat(intVal.(float64), DEFAULT_FLOAT_FMT, DEFAULT_FLOAT_PREC, FLOAT_BITSIZE_64)
	}

	return
}
