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
func boom(cool int64, mool int64) (int64, int64, int64, []int64, []int64) {
	runtime.DfMark(cool) // 1
	runtime.DfMark(mool) // 2
	foo := &Foo{}

	// if cool > 3 {
	// 	foo.bar2[cool] = 20
	// } else {
	// 	foo.bar2[mool] = 40
	// }

	// x := cool

	if cool > 3 {
		foo = &Foo{}
	} else {
		foo = &Foo{bar1: 2.0}
	}

	x := foo.bar2[2]
	y := foo.bar2[4]
	p := runtime.DfInspect(x)
	q := runtime.DfInspect(y)
	dfarr := runtime.DfGetArr()
	blockdfarr := runtime.DfGetBlockArr()
	return p, q, 0, dfarr, blockdfarr
	// return x, y, 0
}

func main() {
	p, q, midx, dfarr, blockdfarr := boom(3, 4)
	fmt.Println(p, q, midx)
	fmt.Println(dfarr)
	fmt.Println(blockdfarr)

	// for idx, val := range dfarr {
	// 	if val > 10 {
	// 		fmt.Println("df of idx")
	// 		fmt.Println(idx, val)
	// 		break
	// 	}
	// }
}
