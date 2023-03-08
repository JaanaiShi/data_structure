package main

import "strconv"

func evalRPN(tokens []string) int {
	var (
		stack  = make([]string, 0)
	)

	for i:=0; i < len(tokens); i++ {
		if tokens[i] != "/" && tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" {
			stack = append(stack, tokens[i])
		}else {
			num1, _ := strconv.Atoi(stack[len(stack) - 1])
			num2, _ := strconv.Atoi(stack[len(stack) - 2])
			stack = stack[:len(stack) - 2]
			result := 0
			if tokens[i] == "/" {
				result = num2 / num1
			} else if tokens[i] == "+" {
				result = num2 + num1
			} else if tokens[i] == "-" {
				result = num2 - num1
			} else if tokens[i] == "*" {
				result = num2 * num1
			}

			stack = append(stack, strconv.Itoa(result))
		}

		
	}
	result, _ := strconv.Atoi(stack[len(stack) -1])
	return result
}


func main() {
	nums := []string{"4","13","5","/","+"}

	evalRPN(nums)
}