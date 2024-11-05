package model_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelFieldListState struct{}

func NewModelFieldListState() *ModelFieldListState {
	return &ModelFieldListState{}
}

func (s *ModelFieldListState) Init(m *configure_command.Model) tea.Cmd {
	items := []list.Item{}
	for _, field := range m.EditingModel.Fields {
		items = append(items, common.Item{Name: field.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Field"
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ModelFieldListState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			if selectedItem, ok := m.List.SelectedItem().(common.Item); ok {
				if selectedItem.Disabled {
					m.Err = fmt.Errorf("This option is not available.")
					return cmd
				}
				for i, field := range m.EditingModel.Fields {
					if field.Name == selectedItem.Name {
						m.EditingFieldIndex = i
						m.EditingField = &m.EditingModel.Fields[i]
						if m.RemovingField {
							// Remove the field
							m.EditingModel.Fields = append(m.EditingModel.Fields[:i], m.EditingModel.Fields[i+1:]...)
							m.RemovingField = false
							m.SetState(NewModelEditState())
						} else {
							// Edit the field
							m.SetState(NewModelEditFieldMenuState())
						}
						break
					}
				}
			}
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelEditState())
		}
	}

	return cmd
}

func (s *ModelFieldListState) View(m *configure_command.Model) string {
	return m.List.View()
}
