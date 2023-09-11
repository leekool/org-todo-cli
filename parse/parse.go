package parse

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
	"regexp"
	"strings"
)

type Todo struct {
	Status string
	Task   string
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

	status := []string{"TODO", "DONE", "SKIP"}

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
			lines = append(lines, todo)
			// styledTodo := styleTodo(todo)
			// lines = append(lines, styledTodo)
		}
	}

	return lines
}

func Toggle(todo Todo) string {
	// style the expected values "TODO" and "DONE"
	styledTodo := styleStatus("TODO")
	styledDone := styleStatus("DONE")

	newStatus := todo.Status

	// compare the styled string with the styled versions of "TODO" and "DONE"
	if todo.Status == styledTodo {
		newStatus = styleStatus("DONE")
	} else if todo.Status == styledDone {
		newStatus = styleStatus("TODO")
	}

	return newStatus
}

func createTodo(input string) Todo {
	parts := strings.SplitN(input, " ", 2)

	t := Todo{
		Status: styleStatus(parts[0]),
		Task:   parts[1],
	}

	return t
}

func styleStatus(input string) string {
	statusTodoStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")) // white
	statusDoneStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2")) // green
	statusSkipStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1")) // maroon

	var styledStatus string

	switch input {
	case "TODO":
		styledStatus = statusTodoStyle.Render(input)
	case "DONE":
		styledStatus = statusDoneStyle.Render(input)
	case "SKIP":
		styledStatus = statusSkipStyle.Render(input)
	}

	return styledStatus
}
