package model_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/states"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelListState struct{}

func NewModelListState() *ModelListState {
	return &ModelListState{}
}

func (s *ModelListState) Init(m *configure_command.Model) tea.Cmd {
	items := []list.Item{}
	for _, model := range m.Data.Models {
		items = append(items, common.Item{Name: model.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Model"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.RemovingModel = false // Reset removing flag

	return nil
}

func (s *ModelListState) Update(msg tea.Msg, m *configure_command.Model) tea.Cmd {
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
				// Find the model index
				for i, model := range m.Data.Models {
					if model.Name == selectedItem.Name {
						if m.RemovingModel {
							// Remove the model
							m.Data.Models = append(m.Data.Models[:i], m.Data.Models[i+1:]...)
							m.RemovingModel = false
							m.SetState(states.NewMainMenuState())
						} else {
							// Edit the model
							m.EditingModelIndex = i
							m.EditingModel = &m.Data.Models[i]
							m.TextInput = textinput.New()
							m.TextInput.Placeholder = "Model Name"
							m.TextInput.SetValue(m.EditingModel.Name)
							m.TextInput.Focus()
							m.SetState(NewModelEditState())
						}
						break
					}
				}
			}
		} else if msg.Type == tea.KeyEsc {
			m.SetState(states.NewMainMenuState())
		}
	}

	return cmd
}

func (s *ModelListState) View(m *configure_command.Model) string {
	return m.List.View()
}
