package model_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelEditFieldJSONNameState struct{}

func NewModelEditFieldJSONNameState() *ModelEditFieldJSONNameState {
	return &ModelEditFieldJSONNameState{}
}

func (s *ModelEditFieldJSONNameState) Init(m *configure_command.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "JSON Name"
	m.TextInput.SetValue(m.EditingField.JSONName)
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ModelEditFieldJSONNameState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			jsonName := m.TextInput.Value()
			if jsonName == "" {
				jsonName = m.EditingField.Name
			} else if jsonName == m.EditingField.JSONName {
				// do nothing, keep the current json's tag and return to the field edit menu
				m.SetState(NewModelEditFieldMenuState())
			}
			m.EditingField.JSONName = jsonName
			m.SetState(NewModelEditFieldMenuState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelEditFieldMenuState())
		}
	}

	return cmd
}

func (s *ModelEditFieldJSONNameState) View(m *configure_command.Model) string {
	return common.RenderPrompt("Enter the new JSON name (leave empty to keep the current json's tag):", m.TextInput.View())
}
