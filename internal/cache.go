package cache

import (
	"fmt"
	"reflect"
)

func Cache[T any](fn any) func(...T) T {
	fnValue := reflect.ValueOf(fn)
	n := fnValue.Type().NumIn()
	fmt.Println(n)
	mp := make(map[string]T)
	switch v := reflect.ValueOf(fn); v.Kind() {
	case reflect.Func:
	default:
		panic("type should be function")
	}

	return func(args ...T) T {
		if len(args) != n {
			panic("invalid number of args")
		}

		key := ToString(args)

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

func ToString[T any](inp []T) string {
	res := ""
	for _, val := range inp {
		res += fmt.Sprintf("%v,", val)
	}
	return res
}
