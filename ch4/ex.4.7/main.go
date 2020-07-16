package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUtf8(bs []byte) []byte {
	reversedSize := 0
	restSize := len(bs)
	for restSize > 0 {
		_, size := utf8.DecodeLastRune(bs)
		bs = append(bs[0:reversedSize], append(bs[len(bs)-size:], bs[reversedSize:len(bs)-size]...)...)
		reversedSize += size
		restSize -= size
	}
	return bs
}

func main() {
	bs := []byte("Hello, 世界")
	bs = reverseUtf8(bs)
	fmt.Println(string(bs))
}
