package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	arguments := strings.Join(args, " ")
	fmt.Print(arguments)
	setup(arguments)
}
