package service_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceListState struct{}

func NewServiceListState() *ServiceListState {
	return &ServiceListState{}
}

func (s *ServiceListState) Init(m *model.Model) tea.Cmd {
	items := []list.Item{}
	for _, svc := range m.Data.Services {
		items = append(items, common.Item{Name: svc.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Service"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.RemovingService = false // Reset removing flag

	return nil
}

func (s *ServiceListState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			if selectedItem, ok := m.List.SelectedItem().(common.Item); ok {
				if selectedItem.Disabled {
					m.Err = fmt.Errorf("this option is not available")
					return cmd
				}
				// Find the service index
				for i, svc := range m.Data.Services {
					if svc.Name == selectedItem.Name {
						if m.RemovingService {
							// Remove the service
							m.Data.Services = append(m.Data.Services[:i], m.Data.Services[i+1:]...)
							m.RemovingService = false
							return func() tea.Msg {
								return model.ReturnToMainMenu{}
							}
						} else {
							// Edit the service
							m.EditingServiceIndex = i
							m.EditingService = &m.Data.Services[i]
							m.TextInput = textinput.New()
							m.TextInput.Placeholder = "Service Name"
							m.TextInput.SetValue(m.EditingService.Name)
							m.TextInput.Focus()
							m.SetState(NewServiceEditMenuState())
						}
						break
					}
				}
			}
		} else if msg.Type == tea.KeyEsc {
			return func() tea.Msg {
				return model.ReturnToMainMenu{}
			}
		}
	}

	return cmd
}

func (s *ServiceListState) View(m *model.Model) string {
	return m.List.View()
}
