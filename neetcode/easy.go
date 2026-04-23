package neetcode

import (
	"container/heap"
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

func TwoSumTest(nums []int, target int) ([]int, bool) {
	totalLen := len(nums) - 1
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	slices.Sort(nums)
	i := 0
	for {
		if i >= totalLen {
			return []int{}, false
		}
		sum := nums[i] + nums[totalLen]
		if sum == target {
			fmt.Println("sum", sum)
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
			return []int{nums[a], nums[b]}, true
		}
		if sum > target {
			totalLen -= 1
		}

		if sum < target {
			i += 1
		}
	}
}
func TwoSum(nums []int, target int) []int {
	totalLen := len(nums) - 1
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	slices.Sort(nums)
	i := 0
	for {
		if i > len(nums)-1 {
			return []int{}
		}
		sum := nums[i] + nums[totalLen]
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

// func TopKFrequent(nums []int, k int) []int {
// 	numCount := map[int]int{}
//
// 	for _, num := range nums {
// 		numCount[num] += 1
// 	}
// 	numCountArray := [][]int{}
// 	for k, v := range numCount {
//
// 		numCountArray = append(numCountArray, []int{k, v})
// 	}
// 	fmt.Println(numCount)
// 	fmt.Println(numCountArray)
// 	slices.SortFunc(numCountArray, func(a, b []int) int {
// 		return b[1] - a[1]
// 	})
// 	fmt.Println(numCountArray)
// 	results := []int{}
//
// 	for i := 0; i < k; i++ {
// 		results = append(results, numCountArray[i][0])
// 	}
// 	return results
//
// }

type Arr [][2]int

func (a Arr) Len() int           { return len(a) }
func (a Arr) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a Arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *Arr) Pop() any {
	old := *a
	last := len(old)
	lastNum := old[last-1]
	*a = old[:last-1]
	return lastNum
}

func (a *Arr) Push(num any) {
	*a = append(*a, num.([2]int))
}
func TopKFrequent(nums []int, k int) []int {
	results := []int{}

	numCount := map[int]int{}

	for _, num := range nums {
		numCount[num] += 1
	}
	MinHeap := &Arr{}
	for key, v := range numCount {

		heap.Push(MinHeap, [2]int{key, v})
		if MinHeap.Len() > k {
			fmt.Println("increased leagth", MinHeap.Len())
			heap.Pop(MinHeap)
		}
	}
	fmt.Println(*MinHeap)
	for _, pair := range *MinHeap {
		results = append(results, pair[0])
	}
	return results

}

func ProductExceptSelf(nums []int) []int {
	pre := []int{}
	post := make([]int, len(nums))
	result := []int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 {

			pre = append(pre, pre[i-1]*nums[i-1])
		} else {
			pre = append(pre, 1)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {

		if i < len(nums)-1 {
			post[i] = nums[i+1] * post[i+1]
		} else {
			post[i] = 1
		}
	}

	for i := 0; i < len(nums); i++ {
		result = append(result, pre[i]*post[i])
	}

	fmt.Println(pre)
	fmt.Println(post)
	return result

}
