package main

//go:noinline
//go:dataflow
func loopy(x int64) int64 {
	for i := 0; i < 10; i++ {
		x++
	}
	return x
}
