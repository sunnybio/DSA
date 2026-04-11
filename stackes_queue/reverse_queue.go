package stackes_queue

import "fmt"

func Reverse() {

	fmt.Println("hello")

}
func Nearest_one(grid [][]int) {

	result := make([][]int, len(grid))
	n := len(grid)
	for i := 0; i < n; i++ {
		result[i] = make([]int, len(grid[i]))
	}

	for i := range n {
		for j := range len(grid[i]) {

			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}

}
