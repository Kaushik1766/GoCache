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
	if fnValue.Kind() != reflect.Func {
		panic("invalid input type")
	}
	return func(args ...any) T {
		if len(args) != n {
			panic("invalid number of args")
		}

		for i := range n {
			fnArgType := fnValue.Type().In(i)
			passedArgType := reflect.TypeOf(args[i])
			if fnArgType != passedArgType {
				panic("type of input parameters don't match type of function parameters")
			}
		}

		key := toString(args)

		if exists, ok := mp[key]; ok {
			return exists
		}

		reflectArgs := make([]reflect.Value, n)

		for i, val := range args {
			reflectArgs[i] = reflect.ValueOf(val)
		}
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
