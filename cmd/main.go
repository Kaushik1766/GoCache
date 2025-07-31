package main

import (
	"fmt"

	cache "github.com/Kaushik1766/GoCache"
)

func main() {
	fn := cache.Cache[int](func(a int, b float64) int {
		return a + int(b)
	})
	fmt.Println("new call: ", fn(3, 2.3))
	fmt.Println("new call: ", fn(2, 2.3))
	fmt.Println("reused call: ", fn(3, 2.3))
}
