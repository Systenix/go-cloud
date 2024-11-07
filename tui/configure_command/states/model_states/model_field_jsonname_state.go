package model_states

import (
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelFieldJSONNameState struct{}

func NewModelFieldJSONNameState() *ModelFieldJSONNameState {
	return &ModelFieldJSONNameState{}
}

func (s *ModelFieldJSONNameState) Init(m *model.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "JSON Name (leave empty to use field name)"
	m.TextInput.Focus()

	return textinput.Blink
}

func (s *ModelFieldJSONNameState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			jsonName := m.TextInput.Value()
			if jsonName == "" {
				jsonName = m.EditingField.Name
			}
			m.EditingField.JSONName = jsonName
			m.EditingModel.Fields = append(m.EditingModel.Fields, *m.EditingField)
			m.EditingField = nil
			m.SetState(NewModelEditState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelFieldTypeState())
		}
	}

	return cmd
}

func (s *ModelFieldJSONNameState) View(m *model.Model) string {
	return common.RenderPrompt("Enter the JSON name (leave empty to use field name):", m.TextInput.View())
}
