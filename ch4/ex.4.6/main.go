package main

import (
	"fmt"
	"unicode"
)

func compressUnicodeSpace(bs []byte) []byte {
	if len(bs) <= 1 {
		return bs
	}
	for i := len(bs) - 1; i > 0; i-- {
		if uint32(bs[i]) <= unicode.MaxLatin1 && bs[i] == bs[i-1] {
			switch bs[i] {
			case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
				bs[i] = ' '
				rest := bs[i:]
				bs = bs[:i-1]
				bs = append(bs, rest...)
			}
		}
	}
	return bs
}
func main() {
	var bs []byte = []byte("Hello,\n\n世界")
	bs = compressUnicodeSpace(bs)
	fmt.Println(string(bs))
}
