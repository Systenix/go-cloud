package model_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelFieldNameState struct{}

func NewModelFieldNameState() *ModelFieldNameState {
	return &ModelFieldNameState{}
}

func (s *ModelFieldNameState) Init(m *configure_command.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "Field Name"
	m.TextInput.SetValue(m.EditingField.Name)
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ModelFieldNameState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			m.EditingField.Name = m.TextInput.Value()
			m.TextInput.Reset()
			m.SetState(NewModelFieldTypeState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelEditState())
		}
	}

	return cmd
}

func (s *ModelFieldNameState) View(m *configure_command.Model) string {
	return common.RenderPrompt("Enter the field name:", m.TextInput.View())
}
