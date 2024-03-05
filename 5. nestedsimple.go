package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	bar1 float64
	bar2 [5]int64
	bar3 int64
}

//go:noinline
//go:dataflow
func boom(cool int64, mool int64) (int64, int64, int64) {
	runtime.DfMark(cool) // 1
	runtime.DfMark(mool) // 2
	foo := &Foo{}

	foo.bar2[2] = cool
	foo.bar2[3] = mool
	foo.bar2[mool] = cool // arr[4] = 3

	x := foo.bar2[2]
	y := foo.bar2[4]
	z := foo.bar2[3]
	p := runtime.DfInspect(x)
	q := runtime.DfInspect(y)
	r := runtime.DfInspect(z)
	return p, q, r
	// return x, y, 0
}

func main() {
	p, q, r := boom(3, 4)
	fmt.Println(p, q, r)
}
