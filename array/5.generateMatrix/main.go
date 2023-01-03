package main

import (
	"fmt"
)

/*
	1. 模拟顺时针画矩阵的过程：
	* 填充上行从左到右
	* 填充右列从上到下
	* 填充下行从右到左
	* 填充左列从下到上

	使用固定规则去遍历，每画一条边都要坚持一致的左闭又开，或者左开右闭
*/

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n

	// 构造二维数组
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++			
		}
		left++
	}
	return matrix 
}

func generateMatrix1(n int) [][]int {
	var (
		top, bottom = 0, n - 1
		left, right = 0, n - 1
		tar = n * n
		num = 1
		matrix = make([][]int, n)
	)

	// 因为n是变量，所以需要构建二维数组
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// 遍历每个数
	for num <= tar {
		// 每走完一条边，都需要将边界点降低, 走边方式为左闭右开
		// top

		if top >= n || bottom < 0 || left >= n || right < 0 {
			break
		}
		
		// 如果n为奇数的话，此处为中间数。
		if num == tar && left == right {
			matrix[left][right] = num 
		}

		for i:=left; i < right; i++ {
			matrix[top][i] = num
			num++
		}

		

		// right
		for i := top; i < bottom; i++ {
			matrix[i][right] = num
			num++
		}
		
		

		// bottom
		for i := right; i > left; i-- {
			matrix[bottom][i] = num
			num++
		}
		
		

		// left
		for i := bottom; i > top; i-- {
			matrix[i][left] = num
			num++
		}
		top++
		right--
		bottom--
		left++
	}

	return matrix
	

}


func main() {
	n := 5
	result := generateMatrix1(n)
	for i:=0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}