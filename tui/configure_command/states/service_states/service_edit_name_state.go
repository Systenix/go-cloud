package service_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceEditNameState struct{}

func NewServiceEditNameState() *ServiceEditNameState {
	return &ServiceEditNameState{}
}

func (s *ServiceEditNameState) Init(m *configure_command.Model) tea.Cmd {
	m.TextInput.Reset()
	m.TextInput.Placeholder = "Service Name"
	m.TextInput.SetValue(m.EditingService.Name)
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ServiceEditNameState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			m.EditingService.Name = m.TextInput.Value()
			m.SetState(NewServiceEditMenuState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewServiceEditMenuState())
		}
	}

	return cmd
}

func (s *ServiceEditNameState) View(m *configure_command.Model) string {
	return common.RenderPrompt("Enter the new service name:", m.TextInput.View())
}
