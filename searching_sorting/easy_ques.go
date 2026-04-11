package searching_sorting

import (
	"fmt"
	"strconv"
	"strings"
)

func CountSquares(n int) int {

	i := 1
	j := 0
	for {
		temp := i * i
		if temp < n {
			j += 1
			i += 1
		} else {
			break
		}

	}
	return j

}

func MajorityElement(arr []int) int {

	majorElementCount := 0
	mElement := -1
	for _, v := range arr {
		if (mElement != v) && (majorElementCount == 0) {
			majorElementCount = 1
			mElement = v
		} else if mElement == v {
			majorElementCount += 1
		} else if mElement != v {
			majorElementCount -= 1
		}

	}
	if majorElementCount == 0 {
		return -1
	}
	return mElement

}

func kthSmallestNumber(arr []int) int {

	return -1
}

type numRange struct {
	lower  int
	higher int
}

func ReadKth() {

	var inputT string
	fmt.Scanf(inputT)
	T, err := strconv.Atoi(inputT)

	if err != nil {

		return
	}
	for i := 0; i < T; i++ {

		var inputNandQ string
		fmt.Scanf(inputNandQ)

		fmtInput := strings.Split(inputNandQ, " ")
		N, err := strconv.Atoi(fmtInput[0])

		if err != nil {
			return
		}
		Q, err := strconv.Atoi(fmtInput[1])
		fmt.Println(Q)
		for j := 0; j < N; j++ {
			var inputRange string

			fmt.Scanf(inputRange)

			fmtRange := strings.Split(inputRange, " ")

			A, err := strconv.Atoi(fmtRange[0])

			B, err := strconv.Atoi(fmtRange[0])

			if err != nil {
				return
			}
			a := numRange{lower: A, higher: B}
			fmt.Println(a)

		}
	}
}
/*----------- Two Sum ------------*/
func twoSum(nums []int, target int) []int {
    for i 
}
