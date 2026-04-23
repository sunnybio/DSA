package neetcode

import (
	"fmt"
	"slices"
)

func IsPalindrome(s string) bool {

	cleanedString := []rune{}
	for _, v := range s {
		if v >= 65 && v <= 90 {

			cleanedString = append(cleanedString, v+32)
		} else if (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			cleanedString = append(cleanedString, v)
		}
	}
	fmt.Println(string(cleanedString))
	i := 0
	j := len(cleanedString) - 1

	for {
		if i > j {
			break
		}
		if cleanedString[i] != cleanedString[j] {
			return false
		}
		i++
		j--
	}
	return true

}

// func ThreeSum(nums []int) [][]int {
//
// 	result := [][]int{}
//
// 	var sortedNums []int
// 	sortedNums = make([]int, len(nums))
//
// 	copy(sortedNums, nums)
// 	slices.Sort(sortedNums)
// 	fmt.Println(sortedNums)
// 	// for i, v := range nums {
// 	//
// 	// }
// 	// for i := 0; i < len(nums); i++ {
// 	// 	for j := i + 1; j < len(nums); j++ {
// 	// 		findNum := 0 - (nums[i] + nums[j])
// 	// 		n3, res := slices.BinarySearch(sortedNums, findNum)
// 	// 		fmt.Println(n3, res, nums[i], nums[j])
// 	// 	}
// 	// }
// 	for i := 0; i < len(sortedNums); i++ {
// 		for {
// 			if i > 0 {
// 				fmt.Println("i", i)
//
// 				if sortedNums[i] != sortedNums[i-1] || i > len(sortedNums)-2 {
// 					fmt.Println("breaking")
// 					break
// 				} else {
// 					i++
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 		targetNum := 0 - sortedNums[i]
// 		fmt.Println("target num", targetNum)
// 		twoSumResult, check := TwoSumTest(sortedNums[i+1:], 0-sortedNums[i])
// 		if check {
// 			result = append(result, append(twoSumResult, sortedNums[i]))
// 		}
// 		fmt.Println(sortedNums[i], twoSumResult)
// 	}
// 	return result
//
// }

func ThreeSum(nums []int) [][]int {

	result := [][]int{}

	var sortedNums []int
	fmt.Println(nums)
	sortedNums = make([]int, len(nums))

	copy(sortedNums, nums)
	slices.Sort(sortedNums)
	fmt.Println(sortedNums)
	for i := 0; i < len(sortedNums); i++ {
		if i > 0 && sortedNums[i] == sortedNums[i-1] {
			continue
		}
		targetNum := 0 - sortedNums[i]
		j := i + 1
		k := len(sortedNums) - 1

		fmt.Println("i", i)
		for {
			fmt.Println(result)
			fmt.Println(i, j, k, targetNum)
			if j >= k {
				break
			}
			sum := sortedNums[j] + sortedNums[k]
			if sum == targetNum {
				result = append(result, []int{sortedNums[i], sortedNums[j], sortedNums[k]})
				for j < k && sortedNums[j] == sortedNums[j+1] {
					j++
				}
				for j < k && sortedNums[k] == sortedNums[k-1] {
					k--
				}
				j++
				k--
			}
			if sum < targetNum {
				j++
			}
			if sum > targetNum {
				k--
			}

		}
	}
	return result

}
