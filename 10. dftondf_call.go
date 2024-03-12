package main

import "fmt"

//go:dataflow
//go:noinline
func boom(x *float64) *float64 {
	y := new(float64)
	*y = (*x) * 2
	z := humm(y)
	// y = 5
	// z = &y
	return &z
}

//go:noinline
func humm(x *float64) float64 {
	y := (*x) * 3
	return y
}

func main() {
	x := 3.0
	pz := boom(&x)
	fmt.Println(*pz)
}
