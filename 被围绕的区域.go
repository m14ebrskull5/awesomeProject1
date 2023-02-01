package main

import "fmt"

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
// 示例 1：
//
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// 示例 2：
//
// 输入：board = [["X"]]
// 输出：[["X"]]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/surrounded-regions
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func solve(board [][]byte) {
	var m = len(board)
	var n = len(board[0])
	var dir = [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	var father = make([]int, m*n+1)
	for i := 0; i < m*n+1; i++ {
		father[i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			for k := 0; k < 4; k++ {
				ni := i + dir[k][0]
				nj := j + dir[k][1]

				if nj < n && nj >= 0 && ni < m && ni >= 0 {
					if 'O' == board[ni][nj] {
						jointSet(father, nums(i, j, m), nums(ni, nj, m))
					}
				} else {
					jointSet(father, nums(i, j, m), m*n)
				}

			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find(father, nums(i, j, m)) != find(father, m*n) && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < m; i++ {
		fmt.Printf("%c\n", board[i])
	}
}
func nums(i int, j int, m int) int {
	return j*m + i
}
func find(father []int, pos int) int {

	fatherItem := father[pos]
	if fatherItem == pos {
		return fatherItem
	}
	father[pos] = find(father, fatherItem)
	return father[pos]
}

func jointSet(father []int, i int, j int) {
	var fi = find(father, i)
	var fj = find(father, j)

	if fi != fj {

		father[fi] = fj
	}
}
func main() {
	var board = [][]byte{
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
	}

	solve(board)
}
