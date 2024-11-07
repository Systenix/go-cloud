package model_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelEditFieldNameState struct{}

func NewModelEditFieldNameState() *ModelEditFieldNameState {
	return &ModelEditFieldNameState{}
}

func (s *ModelEditFieldNameState) Init(m *model.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "Field Name"
	m.TextInput.SetValue(m.EditingField.Name)
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ModelEditFieldNameState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			m.EditingField.Name = m.TextInput.Value()
			m.SetState(NewModelEditFieldMenuState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelEditFieldMenuState())
		}
	}

	return cmd
}

func (s *ModelEditFieldNameState) View(m *model.Model) string {
	return common.RenderPrompt("Enter the new field name:", m.TextInput.View())
}
