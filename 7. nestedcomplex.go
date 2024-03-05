package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	bar1 float64
	bar2 bool
	bar3 int64
	bar4 [5]Bar
}

type Bar struct {
	bar5 Fubar
}

type Fubar struct {
	bar6 [6]float64
}

//go:dataflow
func boom(cool int64) int64 {
	runtime.DfMark(cool)
	foo := &Foo{}

	foo.bar4[3].bar5.bar6[4] = 20

	if cool > 3 {
		foo = &Foo{}
	} else {
		foo = &Foo{bar1: 2.0}
	}

	foo.bar4[cool].bar5.bar6[4] = 20

	return runtime.DfInspect(foo.bar4[cool].bar5.bar6[4])
}

func main() {
	x := boom(4)
	fmt.Println(x)
}
