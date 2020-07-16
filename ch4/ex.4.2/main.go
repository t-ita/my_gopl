package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	input, _ := ioutil.ReadAll(os.Stdin)

	switch arg {
	case "384":
		fmt.Printf("%x\n", sha512.Sum512(input))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(input))
	default:
		fmt.Printf("%x\n", sha256.Sum256(input))
	}
}
