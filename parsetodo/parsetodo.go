package parsetodo

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
	statusTodoStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")) // white
	statusDoneStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2")) // green
	statusSkipStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1")) // maroon

	var styledStatus string

	switch todo.Status {
	case "TODO":
		styledStatus = statusTodoStyle.Render(todo.Status)
	case "DONE":
		styledStatus = statusDoneStyle.Render(todo.Status)
	case "SKIP":
		styledStatus = statusSkipStyle.Render(todo.Status)
	}

	return Todo{
		Status: styledStatus,
		Task:   todo.Task,
	}
}
