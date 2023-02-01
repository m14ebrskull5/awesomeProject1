package main

import (
	"fmt"
	"math"
)

const n = 5

var a = make([][]int, n)

func init() {
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = math.MaxInt
		}
	}
	a[0][0] = 0
	a[0][1] = 10
	a[1][0] = 10
	a[0][3] = 30
	a[3][0] = 30
	a[0][4] = 100
	a[4][0] = 100
	a[1][2] = 50
	a[2][1] = 50
	a[2][3] = 20
	a[3][2] = 20
	a[2][4] = 10
	a[4][2] = 10
	a[3][4] = 60
	a[4][3] = 60
	fmt.Println(a)
}

func main() {
	floyd()
}

func floyd() {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][k] == math.MaxInt || a[k][j] == math.MaxInt {

				} else {
					a[i][j] = min(a[i][j], a[i][k]+a[k][j])
				}
				//a[i][j] = int(math.Min(float64(a[i][j]), float64(a[i][k]+a[k][j])))

				fmt.Println(a[i][j], a[i][j], a[i][k]+a[k][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Println(i, " ", j, ":", a[i][j])
		}
	}
}

func min(i int, f int) int {
	if i == math.MaxInt && f == math.MaxInt {
		return math.MaxInt
	}
	if i < f {
		return i
	} else {
		return f
	}
}
