package generate_command

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
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

func (m *Model) initializeDatabaseList() {
	items := []list.Item{
		item{name: "PostgreSQL"},
		item{name: "MySQL"},
		item{name: "MongoDB"},
		item{name: "Redis"},
	}
	m.List = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.List.Title = "Select a database:"
	m.List.SetSize(35, 15)
	m.List.Select(0)
}

func (m *Model) initializeProtocolList() {
	items := []list.Item{
		item{name: "REST"},
		// Add more protocols if needed
	}
	m.List = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.List.Title = "Select a communication protocol:"
	m.List.SetSize(35, 15)
	m.List.Select(0)
}

func (m *Model) initializeMessageBrokerList() {
	items := []list.Item{
		item{name: "RabbitMQ"},
		item{name: "Kafka"},
		item{name: "NATS"},
	}
	m.List = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.List.Title = "Select a message broker:"
	m.List.SetSize(35, 15)
	m.List.Select(0)
}

func (m *Model) initializeIncludeDocker() {
	m.Cursor = 1 // Default to 'Yes'
}

func (m *Model) initializeIncludeKubernetes() {
	m.Cursor = 1 // Default to 'Yes'
}

func renderPrompt(prompt, inputView string) string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		lipgloss.NewStyle().Bold(true).Render(prompt),
		inputView,
		"(Press Enter to continue)",
	)
}
