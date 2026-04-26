package neetcode

import (
	"fmt"
	"strconv"
)

type stack []string

func (s *stack) pop() (string, error) {

	slice := *s
	if len(slice) == 0 {
		return "", fmt.Errorf("no element in stack")
	}
	lastElement := slice[len(slice)-1]
	*s = slice[:len(slice)-1]
	return lastElement, nil
}
func (s *stack) push(char string) error {

	slice := *s
	slice = append(slice, char)
	*s = slice
	return nil
}
func IsValid(s string) bool {
	charStack := stack{}
	for _, charByte := range s {
		char := string(charByte)
		fmt.Println(char)
		if char == "(" || char == "{" || char == "[" {
			charStack.push(char)
		} else {
			lastChar, err := charStack.pop()
			if err != nil {
				return false
			}
			if char == ")" {
				if lastChar != "(" {
					return false
				}
			}
			if char == "}" {
				if lastChar != "{" {
					return false
				}
			}
			if char == "]" {
				if lastChar != "[" {
					return false
				}
			}
		}
	}
	if len(charStack) == 0 {
		return true
	}
	return false

}

type MinStack struct {
	stackArr []int
	minArr   []int
}

func Constructor() MinStack {

	stack := MinStack{}
	return stack

}

func (this *MinStack) Push(val int) {

	minStack := *this
	minStack.stackArr = append(minStack.stackArr, val)
	if len(minStack.minArr) > 0 {
		minimunElement := minStack.minArr[len(minStack.minArr)-1]
		if val <= minimunElement {
			minStack.minArr = append(minStack.minArr, val)
		}
	} else {
		minStack.minArr = append(minStack.minArr, val)

	}
	fmt.Println(minStack.minArr)

	*this = minStack
}

func (this *MinStack) Pop() {

	slice := *this
	lastElement := slice.stackArr[len(slice.stackArr)-1]
	minimunElement := slice.minArr[len(slice.minArr)-1]
	if lastElement == minimunElement {
		slice.minArr = slice.minArr[:len(slice.minArr)-1]
	}
	slice.stackArr = slice.stackArr[:len(slice.stackArr)-1]
	*this = slice

}

func (this *MinStack) Top() int {

	slice := *this
	lastElement := slice.stackArr[len(slice.stackArr)-1]
	return lastElement
}

func (this *MinStack) GetMin() int {

	minStack := *this
	minimunElement := minStack.minArr[len(minStack.minArr)-1]
	return minimunElement
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func EvalRPN(tokens []string) int {

	tokenStack := stack{}
	for _, value := range tokens {

		fmt.Println(tokenStack)
		if _, err := strconv.Atoi(value); err == nil {

			tokenStack.push(value)
		} else {
			int1, _ := tokenStack.pop()
			int2, _ := tokenStack.pop()
			v1, _ := strconv.Atoi(int1)
			v2, _ := strconv.Atoi(int2)
			if value == "+" {

				tokenStack.push(strconv.Itoa(v2 + v1))
			}

			if value == "-" {

				tokenStack.push(strconv.Itoa(v2 - v1))
			}
			if value == "*" {

				tokenStack.push(strconv.Itoa(v2 * v1))
			}
			if value == "/" {

				fmt.Println(v2 / v1)
				test := v2 / v1
				fmt.Println(test)
				tokenStack.push(strconv.Itoa(test))
			}
		}
	}
	r, _ := tokenStack.pop()
	result, _ := strconv.Atoi(r)
	return result
}
