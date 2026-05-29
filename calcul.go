package main

import "fmt"
import "math"

func main() {
	var a uint64 = 64
	var b int = 8
	var result uint64 = 1
	result = uint64(math.Pow(float64(a), float64(b)))
	fmt.Println(result)
}