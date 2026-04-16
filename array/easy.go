package array

import (
	"fmt"
	"slices"
)

func IsAnagram(s string, t string) bool {
	letterCount1 := map[rune]int{}
	letterCount2 := map[rune]int{}
	if len(s) != len(t) {
		return false
	}
	for index, char := range s {

		letterCount1[char] += 1

		letterCount2[rune(t[index])] += 1
	}
	for key, _ := range letterCount1 {
		count1 := letterCount1[key]
		count2 := letterCount2[key]
		if count1 != count2 {
			return false
		}
	}
	return true
}

func TwoSum(nums []int, target int) []int {
	totalLen := len(nums) - 1
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	slices.Sort(nums)
	fmt.Println(nums)
	fmt.Println(copyNums)
	i := 0
	for {
		fmt.Println(i, totalLen, "test")
		fmt.Println(nums[i], nums[totalLen])
		sum := nums[i] + nums[totalLen]
		fmt.Println(sum)
		if sum == target {
			a, b := 0, 0
			check := 0
			for j := 0; j < len(copyNums); j++ {

				if nums[i] == copyNums[j] && check == 0 {
					check = 1
					a = j
				} else if nums[totalLen] == copyNums[j] {
					b = j
				}
			}
			return []int{a, b}
		}
		if sum > target {
			totalLen -= 1
		}

		if sum < target {
			i += 1
		}
	}
}

func GroupAnagrams(strs []string) [][]string {

	anagramMap := map[string][]string{}
	results := [][]string{}
	for _, s := range strs {
		// flag := true
		// for k := range anagramMap {
		// 	test := IsAnagram(k, s)
		// 	if test {
		// 		anagramMap[k] = append(anagramMap[k], s)
		// 		flag = false
		// 	}
		//
		// }
		// if flag {
		// 	anagramMap[s] = []string{s}
		// }
		b := []rune(s)
		slices.Sort(b)
		re := string(b)
		anagramMap[re] = append(anagramMap[re], s)

	}
	for _, v := range anagramMap {

		results = append(results, v)
	}

	return results
}

func TopKFrequent(nums []int, k int) []int {
	numCount := map[int]int{}

	for _, num := range nums {
		numCount[num] += 1
	}
	numCountArray := [][]int{}
	for k, v := range numCount {

		numCountArray = append(numCountArray, []int{k, v})
	}
	fmt.Println(numCount)
	fmt.Println(numCountArray)
	slices.SortFunc(numCountArray, func(a, b []int) int {
		return b[1] - a[1]
	})
	fmt.Println(numCountArray)
	results := []int{}

	for i := 0; i < k; i++ {
		results = append(results, numCountArray[i][0])
	}
	return results

}
