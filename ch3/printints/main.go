package main

import (
	"bytes"
	"fmt"
)

// intsToString は fmtSprint(values) に似ていますがカンマを追加します
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	println(intsToString([]int{1, 2, 3}))
}
