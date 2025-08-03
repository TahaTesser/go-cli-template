package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppState int

const (
	MenuView AppState = iota
	DetailView
)

type Model struct {
	state    AppState
	cursor   int
	choices  []string
	selected map[int]struct{}
	err      error
}

func InitialModel() Model {
	return Model{
		state:    MenuView,
		cursor:   0,
		choices:  []string{"Option 1", "Option 2", "Option 3", "Option 4", "Quit"},
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}