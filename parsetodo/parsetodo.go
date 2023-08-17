package parsetodo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Todo struct {
	Status string
	Task   string
}

func createTodo(input string) Todo {
	const separator = " "

	parts := strings.SplitN(input, separator, 2)

	t := Todo{
		Status: parts[0],
		Task:	parts[1],
	}

	return t
}

func Parse() []Todo {
	// temporary, will get org folder eventually rather than file
	filePath := "./todotest.org"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return []Todo{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []Todo{}
	}

	status := []string{"TODO", "DONE"}

	// regex := regexp.MustCompile(`\* (TODO|DONE)`)
	regexStatus := `\* (` + strings.Join(status, "|") + `)`
	regex := regexp.MustCompile(regexStatus)

	var lines []Todo

	for scanner.Scan() {
		line := scanner.Text()

		if regex.MatchString(line) {
			line = strings.Replace(line, "* ", "", 1)
			// lines = append(lines, line)

			task := createTodo(line)
			lines = append(lines, task)
		}
	}

	return lines
}
