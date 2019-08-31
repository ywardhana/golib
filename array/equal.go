package array

import "reflect"

func Equal(arrA interface{}, arrB interface{}) bool {
	value := reflect.ValueOf(arrA)
	set := make(map[interface{}]int)
	for i := 0; i < value.Len(); i++ {
		set[value.Index(i).Interface()]++
	}

	value = reflect.ValueOf(arrB)
	for i := 0; i < value.Len(); i++ {
		if set[value.Index(i).Interface()]-1 < 0 {
			return false
		}
		set[value.Index(i).Interface()]--
	}

	sum := 0
	for _, val := range set {
		sum += val
	}

	return sum == 0
}
