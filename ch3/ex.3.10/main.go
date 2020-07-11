package main

import "bytes"

func main() {
	println(comma("100000000"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	countDown := n%3 + 1
	var buf bytes.Buffer
	for i, c := range s {
		countDown--
		if countDown == 0 {
			if i > 0 {
				buf.WriteByte(',')
			}
			countDown = 3
		}
		buf.WriteByte(byte(c))
	}
	return buf.String()
}
