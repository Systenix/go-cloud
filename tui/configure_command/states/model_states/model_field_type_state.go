package model_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelFieldTypeState struct{}

func NewModelFieldTypeState() *ModelFieldTypeState {
	return &ModelFieldTypeState{}
}

func (s *ModelFieldTypeState) Init(m *model.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "string"},
		common.Item{Name: "int"},
		common.Item{Name: "float64"},
		common.Item{Name: "bool"},
		// ... add other basic types ...
	}
	// Add existing model names to allow nested fields
	for _, model := range m.Data.Models {
		items = append(items, common.Item{Name: model.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select Field Type"
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ModelFieldTypeState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
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
				m.EditingField.Type = selectedItem.Name
				m.SetState(NewModelFieldJSONNameState())
			}
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewModelFieldNameState())
		}
	}

	return cmd
}

func (s *ModelFieldTypeState) View(m *model.Model) string {
	return m.List.View()
}
