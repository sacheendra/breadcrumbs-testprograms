package main

import "fmt"

type Foo struct {
	a, b float64
}

//go:noinline
func boom(x *float64) *float64 {
	y := Foo{}
	y.a = (*x) * 2
	z := humm(&y)
	// y = 5
	// z = &y
	return &z
}

//go:dataflow
//go:noinline
func humm(x *Foo) float64 {
	y := (x.a) * 3
	return y
}

func main() {
	x := 3.0
	pz := boom(&x)
	fmt.Println(*pz)
}
