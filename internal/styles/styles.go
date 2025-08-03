package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("#7D56F4")
	secondaryColor = lipgloss.Color("#F25D94")
	accentColor    = lipgloss.Color("#04B575")
	textColor      = lipgloss.Color("#FAFAFA")
	mutedColor     = lipgloss.Color("#626262")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor)

	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(textColor)

	SelectedStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(primaryColor).
			Bold(true)

	CursorStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	CheckedStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	QuitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87")).
			Bold(true)

	HelpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Padding(1, 0)
)