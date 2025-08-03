package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go-cli-template/internal/tui"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-template",
	Short: "A CLI application built with Bubble Tea and Lipgloss",
	Long: `A command-line interface application that demonstrates
the use of Bubble Tea for TUI functionality and Lipgloss for styling.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := tui.InitialModel()
		p := tea.NewProgram(model, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}