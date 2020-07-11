package main

import "strings"

func main() {
	println(basename("a/b/c.go"))
	println(basename("c.d.go"))
	println(basename("abc"))
}

// basename はディレクトリ要素と . 接尾辞を取り除きます
// e.g., a => a, a.go => a,  a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // "/" が見つからなければ -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
