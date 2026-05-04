package neetcode

import (
	"fmt"
	"slices"
)

func Search1(nums []int, target int) int {

	firstElement := 0
	lastElement := len(nums) - 1

	return BS(nums, target, firstElement, lastElement)
}
func BS(nums []int, target int, start int, last int) int {

	if start > last {
		return -1
	}
	mid := (start + last) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return BS(nums, target, start, mid-1)
	} else {
		return BS(nums, target, mid+1, last)
	}
}
func FindPairs(nums []int, k int) int {
	var uniqueMap [][]int
	uniqueMap = make([][]int, len(nums))

	for x := 0; x < len(uniqueMap); x++ {
		uniqueMap[x] = make([]int, len(nums))
	}
	i := 0
	j := len(nums) - 1
	slices.Sort(nums)
	fmt.Println(nums)
	return fp(nums, k, i, j, uniqueMap)

}

func fp(nums []int, k int, start int, end int, uniqueMap [][]int) int {
	sum := nums[start] - nums[end]
	fmt.Println(start, end, sum)

	fmt.Println(uniqueMap)
	if start > end {
		return 0
	}
	if sum == k || sum == -k {

		if uniqueMap[nums[start]][nums[end]] == 1 {

			return 0 + fp(nums, k, start+1, end-1, uniqueMap)
		} else {
			uniqueMap[nums[start]][nums[end]] = 1

			fmt.Println("got sum", start, end, sum)
			return 1 + fp(nums, k, start+1, end-1, uniqueMap)
		}
	}
	if sum < 0 {
		return fp(nums, k, start, end-1, uniqueMap) + fp(nums, k, start+1, end, uniqueMap)

	} else {

		if sum > k {
			return fp(nums, k, start, end-1, uniqueMap)
		} else {
			return fp(nums, k, start+1, end, uniqueMap)
		}

	}

}
func SearchMatrix(matrix [][]int, target int) bool {

	n := len(matrix)
	m := len(matrix[0])

	fmt.Println(n, m)
	fmt.Println(target, matrix)
	return BS1(matrix, target, []int{0, 0}, []int{n - 1, m - 1})
}

func BS1(nums [][]int, target int, start []int, last []int) bool {

	n1, m1 := start[0], start[1]
	n2, m2 := last[0], last[1]
	cols := len(nums[0])
	startIndex := (n1 * cols) + m1
	lastIndex := (n2 * cols) + m2
	midIndex := (startIndex + lastIndex) / 2

	mid1 := midIndex / cols
	mid2 := midIndex - (mid1 * cols)

	// if n1 < 0 || n2 < 0 || m1 < 0 || m2 < 0 {
	// 	return false
	// }
	// if n1 > len(nums)-1 || n2 > len(nums)-1 || m1 > len(nums)-1 || m2 > len(nums)-1 {
	// 	return false
	// }
	if n1 > n2 {
		return false
	}
	if m1 > m2 {
		return false
	}
	fmt.Println(startIndex, lastIndex, midIndex, mid1, mid2)
	mid := nums[mid1][mid2]
	fmt.Println(mid)
	if mid == target {
		return true
	} else if mid > target {
		newEnd := midIndex - 1

		i1 := newEnd / cols
		i2 := newEnd - (i1 * cols)
		fmt.Println(i1, i2)
		return BS1(nums, target, start, []int{i1, i2})
	} else {

		newstart := midIndex + 1
		return BS1(nums, target, []int{newstart / cols, newstart - ((newstart / cols) * cols)}, last)
	}
}

func MinEatingSpeed(piles []int, h int) int {

	maxPile := 0

	for _, v := range piles {
		if v > maxPile {
			maxPile = v
		}
	}
	i := 1
	j := maxPile
	ans := maxPile
	for {
		if i > j {
			break
		}
		mid := (i + j) / 2
		fmt.Println(mid)

		if checkEatingSpeed(piles, h, mid) {
			ans = mid
			j = mid - 1
		} else {
			i = mid + 1

		}
	}
	return ans

}
func checkEatingSpeed(piles []int, h int, k int) bool {

	currHr := 0
	for _, v := range piles {
		if currHr > h {
			return false
		}
		if v%k == 0 {

			currHr = currHr + (v / k)
		} else {

			currHr = currHr + ((v / k) + 1)
		}
	}
	if currHr > h {
		return false
	}
	return true
}

func FindMin(nums []int) int {

	ans := RBS(nums, 0, len(nums)-1)
	if ans == -1 {
		return nums[0]
	}
	return ans
}

func RBS(nums []int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2

	lastEle := nums[len(nums)-1]
	if nums[mid+1] < nums[mid] {
		return nums[mid+1]
	} else if nums[mid] > lastEle {
		return RBS(nums, mid+1, end)
	} else {
		return RBS(nums, start, mid-1)
	}

}

func Search(nums []int, target int) int {

	fmt.Println(nums)
	return rbs1(nums, target, 0, len(nums)-1)
}

// func rbs1(nums []int, target int, start int, end int) int {

// 	if start > end {
// 		return -1
// 	}

// 	mid := (start + end) / 2

// 	fmt.Println(mid, nums[mid])
// 	if nums[mid] == target {
// 		return mid
// 	} else if nums[start] < nums[mid] {
// 		if nums[start] > target && nums[start] > nums[mid] {

// 			return rbs1(nums, target, mid+1, end)
// 		} else {
// 			return rbs1(nums, target, start, mid-1)
// 		}
// 	} else {
// 		if nums[start] < target && nums[start] > nums[mid] {

// 			return rbs1(nums, target, start, mid-1)
// 		} else {
// 			return rbs1(nums, target, mid+1, end)

// 		}
// 	}
// }

func rbs1(nums []int, target int, start int, end int) int {

	if start > end {
		return -1
	}

	mid := (start + end) / 2

	fmt.Println(start, end, mid, nums[mid])
	if nums[mid] == target {
		return mid
	} else if nums[start] <= nums[mid] {
		if target < nums[mid] && nums[start] <= target {

			return rbs1(nums, target, start, mid-1)
		} else {
			return rbs1(nums, target, mid+1, end)
		}
	} else {
		if target > nums[mid] && target < nums[end] {

			return rbs1(nums, target, mid+1, end)
		} else {
			return rbs1(nums, target, start, mid-1)

		}
	}
}
