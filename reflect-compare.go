package main

import (
	"fmt"
	"reflect"
)

// Check if item is i an array for any type using reflection
func containsAny(items interface{}, value interface{}) bool {
	itemsValue := reflect.ValueOf(items)
	if itemsValue.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < itemsValue.Len(); i++ {
		if reflect.DeepEqual(itemsValue.Index(i).Interface(), value) { // value, short for -> reflect.Value(value).Interface()
			fmt.Printf("%s and %s \n", itemsValue.Index(i).Type().String(), reflect.ValueOf(value).Type().String())
			fmt.Printf("%v and %v \n\n", itemsValue.Index(i), reflect.ValueOf(value))
			return true
		}
	}
	return false
}
