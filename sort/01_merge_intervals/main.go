package main

import (
	"fmt"
)

// merge 合并所有重叠的区间，返回不重叠区间的数组。
// 对应 LeetCode 56. Merge Intervals
//
// 示例 1：
//
//	输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//	输出：[[1,6],[8,10],[15,18]]
//	解释：区间 [1,3] 和 [2,6] 重叠，合并为 [1,6]
//
// 示例 2：
//
//	输入：intervals = [[1,4],[4,5]]
//	输出：[[1,5]]
//	解释：区间 [1,4] 和 [4,5] 可被视为重叠区间
//
// 提示：
//   - 1 <= intervals.length <= 10^4
//   - intervals[i].length == 2
//   - 0 <= start_i <= end_i <= 10^4
func merge(intervals [][]int) [][]int {
	// TODO: 请在此实现你的解法
	temp := make([]int, 0)
	for _, interval := range intervals {
		temp = append(temp, interval[0])
	}

}

// -------------------- 以下为辅助函数，无需修改 --------------------

// intervalsEqual 判断两个二维区间切片是否相等。
func intervalsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// 测试用例
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		// 题目示例
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		// 完全不重叠
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		// 全部合并为一个
		{
			[][]int{{1, 4}, {2, 3}},
			[][]int{{1, 4}},
		},
		// 只有一个区间
		{
			[][]int{{1, 10}},
			[][]int{{1, 10}},
		},
		// 乱序输入
		{
			[][]int{{15, 18}, {1, 3}, {2, 6}, {8, 10}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}

	fmt.Println("开始测试 sort/01_merge_intervals...")
	for _, tt := range tests {
		got := merge(tt.intervals)
		if intervalsEqual(got, tt.want) {
			fmt.Printf("PASS: intervals=%v -> %v\n", tt.intervals, got)
		} else {
			fmt.Printf("FAIL: intervals=%v, want=%v, got=%v\n", tt.intervals, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
