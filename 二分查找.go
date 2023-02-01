package main

import "fmt"

func main() {
	x := []int{1, 4, 6, 21, 21, 21, 21, 24, 45, 145}
	u := binarySearch(x, 21, true)
	n := binarySearch(x, 21, false)
	fmt.Println(u, n)
}

func binarySearch(a []int, b int, lower bool) int {
	l := len(a)
	start := 0
	end := l - 1
	for start < end {
		mid := (start + end) / 2
		if lower {
			if a[mid] < b {
				start = mid + 1
			} else {
				end = mid
			}
		} else {
			if a[mid] > b {
				end = mid - 1

			} else {
				start = mid
			}
		}

	}
	if a[start] == b {
		return start
	} else {
		return -1
	}
}
