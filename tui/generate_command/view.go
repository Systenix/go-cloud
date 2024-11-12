package generate_command

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	switch m.State {
	case StateProjectName:
		return renderPrompt("Enter the project name:", m.TextInput.View())

	case StateProjectPath:
		return renderPrompt("Enter the project path (e.g., /github.com/Username):", m.TextInput.View())

	case StateProtocol, StateMessageBroker:
		return m.List.View()

	case StateDone:
		return "Project setup complete. Generating files..."

	default:
		return "Unknown state"
	}
}

func renderPrompt(prompt, inputView string) string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		lipgloss.NewStyle().Bold(true).Render(prompt),
		inputView,
		"(Press Enter to continue)",
	)
}
