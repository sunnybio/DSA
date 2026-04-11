package stackes_queue

import "fmt"

func FindNextSmallest(arr []int) []int {

	result := make([]int, len(arr))

	for i, _ := range arr {
		result[i] = -1
	}

	var stack []int
	for i, v := range arr {

		fmt.Println(v)
		if len(stack) == 0 {
			stack = append(stack, i)
		}
	}
	return arr
}
