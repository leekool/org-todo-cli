package parsetodo

import (
	"github.com/charmbracelet/lipgloss"
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

var statusStyle = lipgloss.NewStyle().Bold(true)

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

			todo := createTodo(line)
			styledTodo := styleTodo(todo)
			lines = append(lines, styledTodo)
		}
	}

	return lines
}

func createTodo(input string) Todo {
	parts := strings.SplitN(input, " ", 2)

	t := Todo{
		Status: parts[0],
		Task:   parts[1],
	}

	return t
}

func styleTodo(todo Todo) Todo {
	styledStatus := statusStyle.Render(todo.Status)

	t := Todo{
		Status: styledStatus,
		Task:   todo.Task,
	}

	return t
}
