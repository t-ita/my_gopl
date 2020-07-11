package main

import "strings"

func main() {
	println(isAnagram("abc", "bca"))
	println(isAnagram("abc", "abx"))
	println(isAnagram("listen", "silent"))
}

func isAnagram(s1 string, s2 string) bool {
	for _, c := range s1 {
		if strings.IndexByte(s2, byte(c)) < 0 {
			return false
		}
	}
	return true
}
