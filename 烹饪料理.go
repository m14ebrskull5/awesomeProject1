package main

import "fmt"

func main() {
	k := 1
	avs(&k)
	fmt.Println(k)
}

func avs(m *int) {
	num := 2
	m = &num
}
