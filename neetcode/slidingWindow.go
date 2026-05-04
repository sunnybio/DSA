package neetcode

import (
	"fmt"
)

func MaxProfit(prices []int) int {
	maxProfit := 0

	i := 0
	j := 0
	for {

		if i >= len(prices)-1 || j >= len(prices)-1 {
			break
		}
		j = i + 1

		for {
			if j > len(prices)-1 {
				break
			}
			fmt.Println(i, j, prices[i], prices[j])

			if prices[j] < prices[i] {
				i = j
				fmt.Println("hehe")
				break
			}
			currSum := prices[j] - prices[i]
			if currSum > maxProfit {
				maxProfit = currSum
			}
			j++
		}

	}
	return maxProfit
}

func LengthOfLongestSubstring(s string) int {
	charIndex := map[byte]int{}

	maxLen := 0

	if len(s) == 1 {
		return 1
	}

	i, j := 0, 0
	fmt.Println(s[0])
	for i < len(s) && j < len(s) {

		for j < len(s) {

			currIndex, ok := charIndex[s[j]]

			fmt.Println(j, i, maxLen, ok)
			if !ok {
				charIndex[s[j]] = j

				maxLen = max(maxLen, j-i+1)
			} else {
				if currIndex <= j && currIndex >= i {
					fmt.Println("here")
					charIndex[s[j]] = j
					i = currIndex + 1

					// maxLen = max(maxLen, j-i+1)
				} else {

					charIndex[s[j]] = j
					maxLen = max(maxLen, j-i+1)
				}

			}
			j++
		}
	}
	return maxLen
}

func CharacterReplacement(s string, k int) int {

}
