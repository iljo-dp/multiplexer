package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	bottomBarStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 2).
		Align(lipgloss.Center)
)

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	ClearTerminal()

	model, err := NewModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing model: %v\n", err)
		os.Exit(1)
	}
	defer model.Close()

	p := tea.NewProgram(model)

	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v\n", err)
		os.Exit(1)
	}
}
