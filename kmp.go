package main

import "fmt"

func main() {
	s := "abc"
	of := indexOf(s, "c")
	fmt.Println(of)
}
func rotateString(s string, goal string) bool {
	s = s + s
	if indexOf(s, goal) != -1 {
		return true
	} else {
		return false
	}
}

func indexOf(str string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	m := len(str)
	n := len(pattern)
	pattern = " " + pattern

	str = " " + str
	next := make([]int, n+1)
	fmt.Println(next)
	for i, j := 2, 0; i < n+1; i++ {
		for j > 0 && pattern[i] != pattern[j+1] {
			j = next[j]
		}
		if pattern[i] == pattern[j+1] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
	for i, j := 1, 0; i < m+1; i++ {
		for j > 0 && str[i] != pattern[j+1] {
			j = next[j]
		}
		if str[i] == pattern[j+1] {
			j++
		}
		fmt.Println(j, n)
		if j == n {
			return i - n
		}
	}
	return -1

}
