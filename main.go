package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of our application.
type Model struct{}

// Init is called when the program starts and returns an optional command.
func (m Model) Init() tea.Cmd {
    return nil
}

// Update is called when a message is received.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        }
    }
    return m, nil
}

// View returns the UI as a string.
func (m Model) View() string {
    return "Hello, Bubble Tea!\nPress q to quit."
}

func main() {
    p := tea.NewProgram(Model{})
    if err := p.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "Error starting program: %v\n", err)
        os.Exit(1)
    }
}
