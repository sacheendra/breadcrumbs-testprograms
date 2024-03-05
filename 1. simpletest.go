package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

//go:noinline
//go:dataflow
func boom(i, j int64) (int64, int64, int64, int64) {
	a := int64(11)
	b := i * a
	c := j * 2
	d := c + 1
	e := c << 3

	runtime.DfMark(d)
	runtime.DfMark(e)
	f := e * d
	g := runtime.DfInspect(f)

	return b, c, d, g
}

//go:noinline
//go:dataflow
func humm(i, j int64) (int64, int64, int64, int64) {
	var x1, x2, x3, x4 int64
	x1, x2, x3, x4 = boom(i, j)
	return x1, x2, x3, x4
}

func main() {
	i, _ := strconv.ParseInt(os.Args[1], 10, 32)
	j, _ := strconv.ParseInt(os.Args[2], 10, 32)
	r1, r2, r3, r4 := humm(i, j)
	fmt.Println(r1, r2, r3, r4)
}
