package configure_command

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().Bold(true).Underline(true)
	errorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	promptStyle       = lipgloss.NewStyle().Bold(true)
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	columnStyle       = lipgloss.NewStyle().Padding(1, 2)
	divider           = lipgloss.NewStyle().
				SetString(" | ").
				Padding(0, 1).
				Foreground(lipgloss.Color("240"))
	disabledItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Italic(true)
	// Add more styles as needed
)
