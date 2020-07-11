package main

import (
	"bytes"
	"strings"
)

func main() {
	println(comma("-100000000.12345"))
}

func comma(s string) string {
	var sign = ""
	var intPart = ""
	var decimalPart = ""

	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		s = s[1:]
	}
	pos := strings.Index(s, ".")
	if pos >= 0 {
		intPart = s[:pos]
		decimalPart = s[pos+1:]
	}
	n := len(intPart)
	if n <= 3 {
		return s
	}
	countDown := n%3 + 1
	var buf bytes.Buffer
	if len(sign) > 0 {
		buf.WriteString(sign)
	}
	for i, c := range intPart {
		countDown--
		if countDown == 0 {
			if i > 0 {
				buf.WriteByte(',')
			}
			countDown = 3
		}
		buf.WriteByte(byte(c))
	}
	if len(decimalPart) > 0 {
		buf.WriteByte('.')
		buf.WriteString(decimalPart)
	}
	return buf.String()
}
