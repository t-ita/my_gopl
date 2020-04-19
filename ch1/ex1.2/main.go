// Echo2 displays command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, " ", arg)
	}
}