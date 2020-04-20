package main

import (
	"strings"
)

func Echo3(args []string) {
	strings.Join(args[1:], " ")
}
