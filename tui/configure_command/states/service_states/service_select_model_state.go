package service_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceSelectModelState struct{}

func NewServiceSelectModelState() *ServiceSelectModelState {
	return &ServiceSelectModelState{}
}

func (s *ServiceSelectModelState) Init(m *model.Model) tea.Cmd {
	items := []list.Item{}
	assignedModels := make(map[string]bool)
	for _, modelName := range m.EditingService.Models {
		assignedModels[modelName] = true
	}

	for _, model := range m.Data.Models {
		items = append(items, common.SelectableItem{
			Name:     model.Name,
			Selected: assignedModels[model.Name],
		})
	}
	if len(items) == 0 {
		m.Err = fmt.Errorf("No models available. Please add models first.")
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = fmt.Sprintf("Assign Models to %s (Space to toggle, Enter to confirm)", m.EditingService.Name)
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ServiceSelectModelState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "space" {
			// Toggle selection
			if item, ok := m.List.Items()[m.List.Index()].(common.SelectableItem); ok {
				item.Selected = !item.Selected
				m.List.SetItem(m.List.Index(), item)
			}
		} else if msg.String() == "enter" {
			// Collect selected models
			m.EditingService.Models = getSelectedModelNames(m.List.Items())
			m.SetState(NewServiceEditMenuState())
		} else if msg.Type == tea.KeyEsc {
			m.SetState(NewServiceEditMenuState())
		}
	}

	return cmd
}

func (s *ServiceSelectModelState) View(m *model.Model) string {
	helpText := "\n\nPress Space to toggle selection, Enter to confirm, Esc to go back."
	return m.List.View() + helpText
}

func getSelectedModelNames(items []list.Item) []string {
	var selectedModels []string
	for _, listItem := range items {
		if item, ok := listItem.(common.SelectableItem); ok && item.Selected {
			selectedModels = append(selectedModels, item.Name)
		}
	}
	return selectedModels
}
