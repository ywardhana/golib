package array

import "reflect"

// Subset check whether the array1 is the subset of array2
func Subset(array1, array2 interface{}) bool {
	firstArray := reflect.ValueOf(array1)
	secondArray := reflect.ValueOf(array2)

	set := make(map[interface{}]int)
	for i := 0; i < secondArray.Len(); i++ {
		set[secondArray.Index(i).Interface()] += 1
	}

	for i := 0; i < firstArray.Len(); i++ {
		value := firstArray.Index(i).Interface()
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
