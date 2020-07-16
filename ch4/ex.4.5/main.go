package main

import "fmt"

func deleteAdjacentDuplication(strs []string) []string {
	if len(strs) <= 1 {
		return strs
	}
	for i := len(strs) - 1; i > 0; i-- {
		if strs[i] == strs[i-1] {
			rest := strs[i:]
			strs = strs[:i-1]
			strs = append(strs, rest...)
		}
	}
	return strs
}

func main() {
	strs := []string{"ABC", "ABC", "DEF", "DEF", "DEF", "EFG"}
	strs = deleteAdjacentDuplication(strs)
	fmt.Println(strs)
}
