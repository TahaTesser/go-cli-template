package tui

import (
	"fmt"

	"go-cli-template/internal/styles"
)

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	s := styles.TitleStyle.Render("Go CLI Template")
	s += "\n\n"

	s += "What would you like to do?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = styles.CursorStyle.Render(">")
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = styles.CheckedStyle.Render("x")
		}

		if i == len(m.choices)-1 {
			s += fmt.Sprintf("%s [%s] %s", cursor, checked, styles.QuitStyle.Render(choice))
		} else {
			line := fmt.Sprintf("%s [%s] %s", cursor, checked, choice)
			if m.cursor == i {
				s += styles.SelectedStyle.Render(line)
			} else {
				s += styles.ItemStyle.Render(line)
			}
		}
		s += "\n"
	}

	s += "\n" + styles.HelpStyle.Render("Press q to quit, ↑/↓ to navigate, space/enter to select")

	return s
}