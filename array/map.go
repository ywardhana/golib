package array

import "reflect"

func Map(slice interface{}, callback func(interface{}) (interface{}, error)) (interface{}, error) {
	values := reflect.ValueOf(slice)
	results := make([]interface{}, 0, values.Len())

	for i := 0; i < values.Len(); i++ {
		result, err := callback(values.Index(i).Interface())

		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
