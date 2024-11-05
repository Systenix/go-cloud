package styles

import "github.com/charmbracelet/lipgloss"

// Styles
var (
	TitleStyle        = lipgloss.NewStyle().Bold(true).Underline(true)
	ErrorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	PromptStyle       = lipgloss.NewStyle().Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ColumnStyle       = lipgloss.NewStyle().Padding(1, 2)
	Divider           = lipgloss.NewStyle().
				SetString(" | ").
				Padding(0, 1).
				Foreground(lipgloss.Color("240"))
	DisabledItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Italic(true)
	// Add more styles as needed
)
