package service_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/states"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceEditMenuState struct{}

func NewServiceEditMenuState() *ServiceEditMenuState {
	return &ServiceEditMenuState{}
}

func (s *ServiceEditMenuState) Init(m *configure_command.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "Edit Service Name"},
		common.Item{Name: "Edit Service Type"},
		common.Item{Name: "Assign Models"},
		common.Item{Name: "Done Editing"},
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = fmt.Sprintf("Editing Service: %s", m.EditingService.Name)
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ServiceEditMenuState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
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
				case "Edit Service Name":
					m.SetState(NewServiceEditNameState())
				case "Edit Service Type":
					m.SetState(NewServiceSelectTypeState())
				case "Assign Models":
					m.SetState(NewServiceSelectModelState())
				case "Done Editing":
					// Save changes and return to main menu
					if m.EditingServiceIndex >= 0 && m.EditingServiceIndex < len(m.Data.Services) {
						m.Data.Services[m.EditingServiceIndex] = *m.EditingService
					} else {
						m.Data.Services = append(m.Data.Services, *m.EditingService)
					}
					m.EditingService = nil
					m.EditingServiceIndex = -1
					m.SetState(states.NewMainMenuState())
				}
			}
		} else if msg.Type == tea.KeyEsc {
			// Go back to main menu
			m.SetState(states.NewMainMenuState())
		}
	}

	return cmd
}

func (s *ServiceEditMenuState) View(m *configure_command.Model) string {
	return m.List.View()
}
