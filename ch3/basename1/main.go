package main

func main() {
	println(basename("a/b/c.go"))
	println(basename("c.d.go"))
	println(basename("abc"))
}

// basename はディレクトリ要素と . 接尾辞を取り除きます
// e.g., a => a, a.go => a,  a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// 最後の'/'とその前のすべてを破棄する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
