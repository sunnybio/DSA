package main

import (
	"fmt"
	"source_code/array"
)

func main() {

	r := array.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(r)
}
