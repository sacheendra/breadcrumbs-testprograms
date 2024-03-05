package main

//go:dataflow
func boom() float64 {
	var arr [5]float64

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// arr[5] = 6

	sum := float64(0)
	for _, num := range arr {
		sum += num
	}

	return sum
}

func hummm() float64 {
	arr := make([]float64, 5)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	sum := float64(0)
	for _, num := range arr {
		sum += num
	}

	return sum
}
