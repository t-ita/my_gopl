package main

import "fmt"

func rotate(s []int, n int) []int {
	if len(s) <= n {
		return s
	}
	s = append(s[n:], s[:n]...)
	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 3)
	fmt.Println(s)
}
