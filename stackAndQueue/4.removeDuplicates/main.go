package main


func removeDuplicates(s string) string {
	var (
		stack = make([]rune, 0)
	)

	for _, v := range s {
		if len(stack) > 0 && stack[len(stack) - 1] == v {
			stack = stack[:len(stack) - 1]
		}else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}


func main() {

}