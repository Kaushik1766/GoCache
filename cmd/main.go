package main

import (
	"cache"
	"fmt"
)

func main() {
	fn := cache.Cache[float64](func(a int, b float64) float64 {
		return float64(a) + b
	})
	fmt.Println("new call: ", fn(3, 2.3))
	fmt.Println("new call: ", fn(2, 2.3))
	fmt.Println("reused call: ", fn(3, 2.3))
}
