package array

import (
	"reflect"
)

func Contains(slice interface{}, elem interface{}) bool {
	value := reflect.ValueOf(slice)

	for i := 0; i < value.Len(); i++ {
		if value.Index(i).Interface() == elem {
			return true
		}
	}
	return false
}
