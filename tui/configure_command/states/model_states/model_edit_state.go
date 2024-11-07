package model_states

import (
	"fmt"

	"github.com/Systenix/go-cloud/generators"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelEditState struct{}

func NewModelEditState() *ModelEditState {
	return &ModelEditState{}
}

func (s *ModelEditState) Init(m *model.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "Add Field"},
		common.Item{Name: "Edit Field"},
		common.Item{Name: "Remove Field"},
		common.Item{Name: "Done Editing"},
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = fmt.Sprintf("Editing Model: %s", m.EditingModel.Name)
	m.List.SetSize(40, 15)
	m.List.Select(0)

	return nil
}

func (s *ModelEditState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd

	if m.EditingModel.Name == "" {
		// Prompt for model name
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.EditingModel.Name = m.TextInput.Value()
				m.SetState(NewModelEditState())
			} else if msg.Type == tea.KeyEsc {
				return func() tea.Msg {
					return model.ReturnToMainMenu{}
				}
			}
		}
	} else {
		// Show field menu
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(common.Item); ok {
					if selectedItem.Disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return cmd
					}
					switch selectedItem.Name {
					case "Add Field":
						m.EditingField = &generators.Field{}
						m.SetState(NewModelFieldNameState())
					case "Edit Field":
						if len(m.EditingModel.Fields) == 0 {
							m.Err = fmt.Errorf("no fields to edit")
						} else {
							m.SetState(NewModelFieldListState())
						}
					case "Remove Field":
						if len(m.EditingModel.Fields) == 0 {
							m.Err = fmt.Errorf("no fields to remove")
						} else {
							m.RemovingField = true
							m.SetState(NewModelFieldListState())
						}
					case "Done Editing":
						// Save or update the model
						if m.EditingModelIndex >= 0 && m.EditingModelIndex < len(m.Data.Models) {
							m.Data.Models[m.EditingModelIndex] = *m.EditingModel
						} else {
							m.Data.Models = append(m.Data.Models, *m.EditingModel)
						}
						m.EditingModel = nil
						m.EditingModelIndex = -1
						return func() tea.Msg {
							return model.ReturnToMainMenu{}
						}
					}
				}
			} else if msg.Type == tea.KeyEsc {
				return func() tea.Msg {
					return model.ReturnToMainMenu{}
				}
			}
		}
	}

	return cmd
}

func (s *ModelEditState) View(m *model.Model) string {
	return m.List.View()
}
