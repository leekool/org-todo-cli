package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"org-todo-cli/parsetodo"
	"os"
)

type todo struct {
	status string
	task   string
}

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func getTasks() []string {
	tasks := parsetodo.Parse()

	var combinedTasks []string

	for _, t := range tasks {
		combinedTask := t.Status + " " + t.Task
		combinedTasks = append(combinedTasks, combinedTask)
	}

	return combinedTasks
}

func initialModel() model {
	return model{
		choices:  getTasks(),
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// return 'nil' (no i/o)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// key press
	case tea.KeyMsg:

		// actual key pressed
		switch msg.String() {

		// exit keys
		case "ctrl+c", "q":
			return m, tea.Quit

		// move cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// move cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// the "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// return updated model to bubble tea runtime for processing
	// note that we're not returning a command
	return m, nil
}

func (m model) View() string {
	var s string

	// iterate over choices
	for i, choice := range m.choices {

		// is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		// is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected
		}

		// render row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// footer
	s += "\n(q) quit.\n"

	// send UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
