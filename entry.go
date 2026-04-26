package main

import (
	"fmt"
	"source_code/neetcode"
)

func main() {

	r := neetcode.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	fmt.Println(r)
}
