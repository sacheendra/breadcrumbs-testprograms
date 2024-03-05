package main

import "fmt"

//go:noinline
//go:dataflow
func boom(x int64) int64 {
	y := x + 1
	return y
}

func main() {
	x := int64(3)
	y := boom(x)
	fmt.Println(y)
}
