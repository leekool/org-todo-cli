package main

import (
	"bufio"
	"fmt"
	"regexp"
	// "io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// temporary, will get org folder eventually rather than file
	filePath := "./todotest.org"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile(`\* (TODO|DONE)`)

	for scanner.Scan() {
		line := scanner.Text()

		if regex.MatchString(line) {
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
