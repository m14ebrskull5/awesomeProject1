package main

import "fmt"

func main() {
	times := [][]int{
		{1, 2, 1},
		{2, 3, 7},
		{1, 3, 4},
		{2, 1, 2},
	}
	time := networkDelayTime(times, 3, 2)
	fmt.Println(time)
}

var dist []int

func networkDelayTime(times [][]int, n int, k int) int {
	dist = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dist[i] = 1000
	}
	dist[k] = 0
	for i := 0; i < n-1; i++ {
		flag := false
		for _, item := range times {

			if dist[item[0]]+item[2] < dist[item[1]] {
				// fmt.Println(dist[item[0]]+ item[2], dist[item[1]])
				dist[item[1]] = dist[item[0]] + item[2]
				flag = true
			}
			fmt.Println(dist[1:])
		}
		if flag == false {
			break
		}

	}
	///[1000 0 1000]
	//[1000 0 7]
	//[1000 0 7]
	//[2 0 7]
	//[2 0 7]
	//[2 0 7]
	//[2 0 6]
	//[2 0 6]
	max := -100
	for i := 1; i <= n; i++ {
		if dist[i] > max {
			max = dist[i]
		}
	}
	if max == 1000 {
		return -1
	} else {
		return max
	}
}
