package main

import (
	cache "cache/internal"
	"fmt"
)

func main() {
	fn := cache.Cache[int](func(a, b int) int {
		return a + b
	})
	fmt.Println(fn(2, 3))
	fmt.Println(fn(2, 3))
}
