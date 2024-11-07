package states

import (
	"fmt"

	"github.com/Systenix/go-cloud/generators"
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/Systenix/go-cloud/tui/configure_command/states/model_states"
	"github.com/Systenix/go-cloud/tui/configure_command/states/service_states"
	"github.com/Systenix/go-cloud/tui/configure_command/styles"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MainMenuState struct{}

func NewMainMenuState() *MainMenuState {
	return &MainMenuState{}
}

func (s *MainMenuState) Init(m *model.Model) tea.Cmd {
	items := []list.Item{
		common.Item{Name: "Add Service"},
		common.Item{Name: "Edit Service", Disabled: len(m.Data.Services) == 0},
		common.Item{Name: "Remove Service", Disabled: len(m.Data.Services) == 0},
		common.Item{Name: "Add Model"},
		common.Item{Name: "Edit Model", Disabled: len(m.Data.Models) == 0},
		common.Item{Name: "Remove Model", Disabled: len(m.Data.Models) == 0},
		common.Item{Name: "Save and Exit"},
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Blueprint Configuration"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.Err = nil // Reset any errors
	return nil
}

func (s *MainMenuState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
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
				case "Add Service":
					m.EditingService = &generators.Service{}
					m.EditingServiceIndex = -1
					m.SetState(service_states.NewServiceEditState())
				case "Edit Service":
					if len(m.Data.Services) == 0 {
						m.Err = fmt.Errorf("no services to edit")
					} else {
						m.SetState(service_states.NewServiceListState())
					}
				case "Remove Service":
					if len(m.Data.Services) == 0 {
						m.Err = fmt.Errorf("no services to remove")
					} else {
						m.SetState(service_states.NewServiceListState())
						m.RemovingService = true
					}

				// Model related
				case "Add Model":
					m.EditingModel = &generators.Model{}
					m.EditingModelIndex = -1
					m.TextInput = textinput.New()
					m.TextInput.Placeholder = "Model Name"
					m.TextInput.Focus()
					m.SetState(model_states.NewModelEditState())
				case "Edit Model":
					if len(m.Data.Models) == 0 {
						m.Err = fmt.Errorf("no models to edit")
					} else {
						m.SetState(model_states.NewModelListState())
					}
				case "Remove Model":
					if len(m.Data.Models) == 0 {
						m.Err = fmt.Errorf("no models to remove")
					} else {
						m.SetState(model_states.NewModelListState())
						m.RemovingModel = true
					}
				case "Save and Exit":
					m.SetState(NewCompleteState())
				}
			}
		}
	}
	return cmd
}

func (s *MainMenuState) View(m *model.Model) string {
	if m.Err != nil {
		errorMsg := styles.ErrorStyle.Render(m.Err.Error())
		mainContent := fmt.Sprintf("%s\n\n%s", m.List.View(), errorMsg)
		m.Err = nil // Reset the error after displaying
		return mainContent
	}
	return m.List.View()
}
