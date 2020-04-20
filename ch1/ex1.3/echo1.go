// Echo1 displays command line arguments.
package main

import (
	"os"
)

func Echo1(args []string) {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}
