package states

import (
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	tea "github.com/charmbracelet/bubbletea"
)

type CompleteState struct{}

func NewCompleteState() *CompleteState {
	return &CompleteState{}
}

func (s *CompleteState) Init(m *model.Model) tea.Cmd {
	return nil
}

func (s *CompleteState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	return tea.Quit
}

func (s *CompleteState) View(m *model.Model) string {
	return "Configuration complete. Saving..."
}
