package parsetodo

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse() []string {
	// temporary, will get org folder eventually rather than file
	filePath := "./todotest.org"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}

	regex := regexp.MustCompile(`\* (TODO|DONE)`)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if regex.MatchString(line) {
			line = strings.Replace(line, "* ", "", 1)
			lines = append(lines, line)
		}
	}

	return lines
}
