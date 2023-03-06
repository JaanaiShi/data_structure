package main

import (
	"fmt"
)

func isValid(s string) bool {
	var (
		leftStack []rune
	)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			leftStack = append(leftStack, v)
		} else {
			length := len(leftStack)

			if length == 0 {
				return false
			} 
			if v == ')' && leftStack[length - 1] == '(' {
				leftStack = leftStack[:length-1]
			} else if v == ']' && leftStack[length - 1] == '[' {
				leftStack = leftStack[:length-1]
			} else if v == '}' && leftStack[length - 1] == '{' {
				leftStack = leftStack[:length-1]
			}else {
				return false
			}
		}
		
	}

	if len(leftStack) != 0 {
		return false
	}

	return true

}

func main() {
	fmt.Println(isValid("([{}])"))

}
