package service_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceSelectTypeState struct{}

func NewServiceSelectTypeState() *ServiceSelectTypeState {
	return &ServiceSelectTypeState{}
}

func (s *ServiceSelectTypeState) Init(m *configure_command.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "rest"},
		// Add other service types as needed
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select Service Type"
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ServiceSelectTypeState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
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
				m.EditingService.Type = selectedItem.Name
				m.SetState(NewServiceEditMenuState())
			}
		} else if msg.Type == tea.KeyEsc {
			// Go back to service edit menu
			m.SetState(NewServiceEditMenuState())
		}
	}

	return cmd
}

func (s *ServiceSelectTypeState) View(m *configure_command.Model) string {
	return m.List.View()
}
