package main

import (
	"fmt"

	cache "github.com/Kaushik1766/GoCache"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fn := cache.Cache[int](func(a int, b float32) int {
		return a + int(b)
	})
	fmt.Println("new call: ", fn(3, float32(2.3)))
	fmt.Println("new call: ", fn(2, float32(2.3)))
	fmt.Println("reused call: ", fn(3.3, float32(2.3)))
}
