package main

import (
	"fmt"
	"leetcode/binaryTree/21.radix/radix"
)


func main() {
	root := radix.NewRadixTree()
	
	
	root.Insert("abc")
	root.Insert("acd")
	root.Insert("ace")
	fmt.Println(root.String())
	
}