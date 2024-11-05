package states

import (
	"github.com/Systenix/go-cloud/tui/configure_command"
	tea "github.com/charmbracelet/bubbletea"
)

type CompleteState struct{}

func NewCompleteState() *CompleteState {
	return &CompleteState{}
}

func (s *CompleteState) Init(m *configure_command.Model) tea.Cmd {
	return nil
}

func (s *CompleteState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
	return tea.Quit
}

func (s *CompleteState) View(m *configure_command.Model) string {
	return "Configuration complete. Saving..."
}
