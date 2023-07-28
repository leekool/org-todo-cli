package main

import (
	// "bufio"
	"fmt"
	// "io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/home/lee/sync/org/todo.org")
	check(err)
	fmt.Print(string(dat))
}
