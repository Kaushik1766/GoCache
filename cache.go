// Package cache provides Cache method to cache function output
package cache

import (
	"fmt"
	"reflect"
)

// Cache takes in an input function and caches its output
func Cache[T any](fn any) func(...any) T {
	fnValue := reflect.ValueOf(fn)
	n := fnValue.Type().NumIn()
	fmt.Println(n)
	mp := make(map[string]T)
	switch v := reflect.ValueOf(fn); v.Kind() {
	case reflect.Func:
	default:
		panic("type should be function")
	}

	return func(args ...any) T {
		if len(args) != n {
			panic("invalid number of args")
		}

		key := toString(args)

		// dp se nikalne ka kam
		if exists, ok := mp[key]; ok {
			fmt.Println("mil gya")
			return exists
		}

		reflectArgs := make([]reflect.Value, n)

		for i, val := range args {
			reflectArgs[i] = reflect.ValueOf(val)
		}
		fmt.Println(reflectArgs)
		result := fnValue.Call(reflectArgs)
		out := result[0].Interface().(T)
		mp[key] = out
		return out
	}
}

func toString[T any](inp []T) string {
	res := ""
	for _, val := range inp {
		res += fmt.Sprintf("%v,", val)
	}
	return res
}
