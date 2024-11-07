package service_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceEditState struct{}

func NewServiceEditState() *ServiceEditState {
	return &ServiceEditState{}
}

func (s *ServiceEditState) Init(m *model.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "Service Name"
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ServiceEditState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	if m.EditingService.Name == "" {
		// We're adding a new service
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.EditingService.Name = m.TextInput.Value()
				m.TextInput.Reset()
				// Transition to service type selection
				m.SetState(NewServiceSelectTypeState())
			} else if msg.Type == tea.KeyEsc {
				// Go back to main menu
				return func() tea.Msg {
					return model.ReturnToMainMenu{}
				}
			}
		}
	} else {
		// We're editing an existing service
		// Transition to service edit menu
		m.SetState(NewServiceEditMenuState())
	}
	return cmd
}

func (s *ServiceEditState) View(m *model.Model) string {
	if m.EditingService.Name == "" {
		return common.RenderPrompt("Enter the service name:", m.TextInput.View())
	} else if m.EditingService.Type == "" {
		return common.RenderPrompt("Enter the service type (e.g., rest):", m.TextInput.View())
	}
	return m.List.View()
}
