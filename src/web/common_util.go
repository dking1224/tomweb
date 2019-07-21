package web

import "reflect"

func StringIsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
