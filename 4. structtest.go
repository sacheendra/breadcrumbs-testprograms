package main

import (
	"fmt"
)

type Foo struct {
	bar1 int64
	bar2 bool
	bar3 int64
	bar4 float64
	bar5 float64
}

//go:noinline
// func boom(x *float64) *float64 {
// 	foo := Foo{}
// 	j := float64(23)
// 	z := &j
// 	j = 99

// 	var cool [5]float64
// 	l := &cool[4]

// 	a := foo.bar1 * int64(*l)
// 	var b float64
// 	if foo.bar2 {
// 		b = float64(foo.bar3 * a)
// 	} else {
// 		b = foo.bar4
// 	}
// 	i := humm(z, l)
// 	i = i * (*x)
// 	foo.bar5 = b * i

// 	y := b * float64(foo.bar5)
// 	// y := 5.0

// 	return &y
// }

//go:noinline
//go:dataflow
func boom(x *float64, y *float64, z float64) (*float64, int64) {
	// if x == nil {
	// 	x = new(float64)
	// }
	// z := x
	a := *x * *y * z
	b := a * z
	c := &b
	return c, 0
}

func main() {
	x := 1.0
	y := 2.0
	z := 3.0
	c, df := boom(&x, &y, z)
	fmt.Println(*c, df)
}
