package stackes_queue

import (
	"fmt"
)

type Stack struct {
	s []int
}

func (s *Stack) Push(v int) []int {
	return append(s.s, v)
}

func (s *Stack) Pop() ([]int, int, error) {

	if len(s.s) == 0 {
		return s.s, -1, fmt.Errorf("empty stack")
	}
	l := len(s.s)
	return s.s[:l-1], s.s[l-1], nil
}

func StackPermutation(s *Stack) *Stack {

	return s
}

func Circulartour(gas []int, cost []int) (int, error) {

	tank := 0
	var result int
	numberOfGas := len(gas)
	flag1 := 0
	i := 0
	for flag1 != 1 {

		flag := 0

		currIndex := i
		j := i
		fmt.Println("i and j", i, j, currIndex)
		for j != currIndex-1 {

			if j >= numberOfGas {
				j = 0
			}
			if tank < 0 {
				flag = 1
				break

			}
			tank += gas[j]
			tank -= cost[j]
			fmt.Println("tank ", tank)
			j += 1
			i += 1
			fmt.Println(tank)
		}
		if flag == 0 {
			result = i
			break
		}

	}
	return result, nil
}
