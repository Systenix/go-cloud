package model_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelEditFieldMenuState struct{}

func NewModelEditFieldMenuState() *ModelEditFieldMenuState {
	return &ModelEditFieldMenuState{}
}

func (s *ModelEditFieldMenuState) Init(m *configure_command.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "Edit Field Name"},
		common.Item{Name: "Edit Field Type"},
		common.Item{Name: "Edit JSON Name"},
		common.Item{Name: "Done Editing Field"},
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = fmt.Sprintf("Editing Field: %s", m.EditingField.Name)
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ModelEditFieldMenuState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
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
				switch selectedItem.Name {
				case "Edit Field Name":
					m.SetState(NewModelEditFieldNameState())
				case "Edit Field Type":
					m.SetState(NewModelEditFieldTypeState())
				case "Edit JSON Name":
					m.SetState(NewModelEditFieldJSONNameState())
				case "Done Editing Field":
					m.EditingField = nil
					m.EditingFieldIndex = -1
					m.SetState(NewModelEditState())
				}
			}
		} else if msg.Type == tea.KeyEsc {
			m.EditingField = nil
			m.EditingFieldIndex = -1
			m.SetState(NewModelEditState())
		}
	}

	return cmd
}

func (s *ModelEditFieldMenuState) View(m *configure_command.Model) string {
	return m.List.View()
}
