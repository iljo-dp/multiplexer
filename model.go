package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	term   *Terminal
	width  int
	height int
	buffer []byte
}

func NewModel() (*Model, error) {
	term, err := NewTerminal()
	if err != nil {
		return nil, err
	}

	return &Model{
		term:   term,
		buffer: make([]byte, 1024),
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			_, _ = m.term.Write([]byte(msg.String()))
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.term.Resize(msg.Width, msg.Height)
	}
	return m, nil
}

func (m *Model) View() string {
	n, _ := m.term.Read(m.buffer)
	return string(m.buffer[:n])
}

func (m *Model) Close() error {
	return m.term.Close()
}
