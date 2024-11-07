package configure_command

import (
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/Systenix/go-cloud/tui/configure_command/states"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *ConfigureModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle state transition messages.
	switch msg.(type) {
	case model.ReturnToMainMenu:
		m.SetState(states.NewMainMenuState())
		return m, nil
	}

	// Current state handle the message
	if m.CurrentState != nil {
		cmd := m.CurrentState.Update(msg, m.Model)
		return m, cmd
	}

	return m, nil
}

/* func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.State {
	case StateProjectInfo:
		m.TextInput, cmd = m.TextInput.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				if m.Data.ProjectName == "" {
					m.Data.ProjectName = m.TextInput.Value()
					m.TextInput.Reset()
					m.TextInput.Placeholder = "Module Path (e.g., github.com/Username/project)"
				} else if m.Data.ModulePath == "" {
					m.Data.ModulePath = m.TextInput.Value()
					m.State = StateMainMenu
					m.initializeMainMenu()
				}
			}
		}

	case StateMainMenu:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					switch selectedItem.name {

					// Service related
					case "Add Service":
						m.State = StateServiceEdit
						m.editingService = &generators.Service{}
						m.editingServiceIndex = -1
						m.TextInput = textinput.New()
						m.TextInput.Placeholder = "Service Name"
						m.TextInput.Focus()
					case "Edit Service":
						if len(m.Data.Services) == 0 {
							m.Err = fmt.Errorf("no services to edit")
						} else {
							m.State = StateServiceList
							m.initializeServiceList()
						}
					case "Remove Service":
						if len(m.Data.Services) == 0 {
							m.Err = fmt.Errorf("no services to remove")
						} else {
							m.State = StateServiceList
							m.initializeServiceList()
							m.removingService = true
						}

					// Model related
					case "Add Model":
						m.State = StateModelEdit
						m.editingModel = &generators.Model{}
						m.editingModelIndex = -1
						m.TextInput = textinput.New()
						m.TextInput.Placeholder = "Model Name"
						m.TextInput.Focus()
					case "Edit Model":
						if len(m.Data.Models) == 0 {
							m.Err = fmt.Errorf("no models to edit")
						} else {
							m.State = StateModelList
							m.initializeModelList()
						}
					case "Remove Model":
						if len(m.Data.Models) == 0 {
							m.Err = fmt.Errorf("no models to remove")
						} else {
							m.State = StateModelList
							m.initializeModelList()
							m.removingModel = true
						}

					case "Save and Exit":
						m.State = StateDone
						return m, tea.Quit
					}
				}
			}
		}

	case StateServiceEdit:
		if m.editingService.Name == "" {
			// We're adding a new service
			m.TextInput, cmd = m.TextInput.Update(msg)
			switch msg := msg.(type) {
			case tea.KeyMsg:
				if msg.Type == tea.KeyEnter {
					m.editingService.Name = m.TextInput.Value()
					m.TextInput.Reset()
					// Transition to service type selection
					m.State = StateSelectServiceType
					m.initializeServiceTypeList()
				} else if msg.Type == tea.KeyEsc {
					// Go back to main menu
					m.State = StateMainMenu
					m.initializeMainMenu()
				}
			}
		} else {
			// We're editing an existing service
			// Transition to service edit menu
			m.State = StateServiceEditMenu
			m.initializeServiceEditMenu()
		}

	case StateServiceEditMenu:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					switch selectedItem.name {
					case "Edit Service Name":
						m.State = StateEditServiceName
						m.TextInput.Reset()
						m.TextInput.Placeholder = "Service Name"
						m.TextInput.SetValue(m.editingService.Name)
						m.TextInput.Focus()
					case "Edit Service Type":
						m.State = StateSelectServiceType
						m.initializeServiceTypeList()
					case "Assign Models":
						m.State = StateSelectModelsForService
						m.initializeModelSelectionList()
					case "Done Editing":
						// Save changes and return to main menu
						if m.editingServiceIndex >= 0 && m.editingServiceIndex < len(m.Data.Services) {
							m.Data.Services[m.editingServiceIndex] = *m.editingService
						} else {
							m.Data.Services = append(m.Data.Services, *m.editingService)
						}
						m.editingService = nil
						m.editingServiceIndex = -1
						m.State = StateMainMenu
						m.initializeMainMenu()
					}
				}
			} else if msg.Type == tea.KeyEsc {
				// Go back to main menu
				m.State = StateMainMenu
				m.initializeMainMenu()
			}
		}

	case StateEditServiceName:
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.editingService.Name = m.TextInput.Value()
				m.State = StateServiceEditMenu
				m.initializeServiceEditMenu()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateServiceEditMenu
				m.initializeServiceEditMenu()
			}
		}

	case StateSelectServiceType:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					m.editingService.Type = selectedItem.name
					m.State = StateServiceEditMenu
					m.initializeServiceEditMenu()
				}
			} else if msg.Type == tea.KeyEsc {
				// Go back to service edit menu
				m.State = StateServiceEditMenu
				m.initializeServiceEditMenu()
			}
		}

	case StateSelectModelsForService:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "space" {
				// Toggle selection
				if item, ok := m.List.Items()[m.List.Index()].(selectableItem); ok {
					item.selected = !item.selected
					m.List.SetItem(m.List.Index(), item)
				}
			} else if msg.String() == "enter" {
				// Collect selected models
				m.editingService.ModelNames = getSelectedModelNames(m.List.Items())
				m.State = StateServiceEditMenu
				m.initializeServiceEditMenu()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateServiceEditMenu
				m.initializeServiceEditMenu()
			}
		}

	case StateServiceList:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("this option is not available")
						return m, cmd
					}
					// Find the service index
					for i, svc := range m.Data.Services {
						if svc.Name == selectedItem.name {
							if m.removingService {
								// Remove the service
								m.Data.Services = append(m.Data.Services[:i], m.Data.Services[i+1:]...)
								m.removingService = false
								m.State = StateMainMenu
								m.initializeMainMenu()
							} else {
								// Edit the service
								m.editingServiceIndex = i
								m.editingService = &m.Data.Services[i]
								m.TextInput = textinput.New()
								m.TextInput.Placeholder = "Service Name"
								m.TextInput.SetValue(m.editingService.Name)
								m.TextInput.Focus()
								m.State = StateServiceEditMenu
								m.initializeServiceEditMenu()
							}
							break
						}
					}
				}
			} else if msg.Type == tea.KeyEsc {
				m.State = StateMainMenu
				m.initializeMainMenu()
			}
		}

	case StateModelEdit:
		if m.editingModel.Name == "" {
			// Prompt for model name
			m.TextInput, cmd = m.TextInput.Update(msg)
			switch msg := msg.(type) {
			case tea.KeyMsg:
				if msg.Type == tea.KeyEnter {
					m.editingModel.Name = m.TextInput.Value()
					m.TextInput.Reset()
					m.initializeFieldMenu()
				} else if msg.Type == tea.KeyEsc {
					m.State = StateMainMenu
					m.initializeMainMenu()
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
					if selectedItem, ok := m.List.SelectedItem().(item); ok {
						if selectedItem.disabled {
							m.Err = fmt.Errorf("This option is not available.")
							return m, cmd
						}
						switch selectedItem.name {
						case "Add Field":
							m.editingField = &generators.Field{}
							m.State = StateFieldName
							m.TextInput.Reset()
							m.TextInput.Placeholder = "Field Name"
							m.TextInput.Focus()
						case "Edit Field":
							if len(m.editingModel.Fields) == 0 {
								m.Err = fmt.Errorf("no fields to edit")
							} else {
								m.State = StateFieldList
								m.initializeFieldList()
							}
						case "Remove Field":
							if len(m.editingModel.Fields) == 0 {
								m.Err = fmt.Errorf("no fields to remove")
							} else {
								m.State = StateFieldList
								m.removingField = true
								m.initializeFieldList()
							}
						case "Done Editing":
							// Save or update the model
							if m.editingModelIndex >= 0 && m.editingModelIndex < len(m.Data.Models) {
								m.Data.Models[m.editingModelIndex] = *m.editingModel
							} else {
								m.Data.Models = append(m.Data.Models, *m.editingModel)
							}
							m.editingModel = nil
							m.editingModelIndex = -1
							m.State = StateMainMenu
							m.initializeMainMenu()
						}
					}
				} else if msg.Type == tea.KeyEsc {
					m.State = StateMainMenu
					m.initializeMainMenu()
				}
			}
		}

	case StateFieldEditMenu:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					switch selectedItem.name {
					case "Edit Field Name":
						m.State = StateEditFieldName
						m.TextInput.Reset()
						m.TextInput.Placeholder = "Field Name"
						m.TextInput.SetValue(m.editingField.Name)
						m.TextInput.Focus()
					case "Edit Field Type":
						m.State = StateEditFieldType
						m.initializeFieldTypeList()
					case "Edit JSON Name":
						m.State = StateEditFieldJSONName
						m.TextInput.Reset()
						m.TextInput.Placeholder = "JSON Name"
						m.TextInput.SetValue(m.editingField.JSONName)
						m.TextInput.Focus()
					case "Done Editing Field":
						m.editingField = nil
						m.editingFieldIndex = -1
						m.State = StateModelEdit
						m.initializeFieldMenu()
					}
				}
			} else if msg.Type == tea.KeyEsc {
				m.editingField = nil
				m.editingFieldIndex = -1
				m.State = StateModelEdit
				m.initializeFieldMenu()
			}
		}

	case StateEditFieldName:
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.editingField.Name = m.TextInput.Value()
				m.State = StateFieldEditMenu
				m.initializeFieldEditMenu()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateFieldEditMenu
				m.initializeFieldEditMenu()
			}
		}

	case StateFieldName:
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.editingField.Name = m.TextInput.Value()
				m.TextInput.Reset()
				m.State = StateFieldType
				m.initializeFieldTypeList()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateModelEdit
				m.initializeFieldMenu()
			}
		}

	case StateEditFieldType:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					m.editingField.Type = selectedItem.name
					m.State = StateFieldEditMenu
					m.initializeFieldEditMenu()
				}
			} else if msg.Type == tea.KeyEsc {
				m.State = StateFieldEditMenu
				m.initializeFieldEditMenu()
			}
		}

	case StateFieldType:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					m.editingField.Type = selectedItem.name
					m.State = StateFieldJSONName
					m.TextInput.Reset()
					m.TextInput.Placeholder = "JSON Name (leave empty to use field name)"
					m.TextInput.Focus()
				}
			} else if msg.Type == tea.KeyEsc {
				m.State = StateFieldName
				m.TextInput.SetValue(m.editingField.Name)
				m.TextInput.Focus()
			}
		}

	case StateEditFieldJSONName:
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				jsonName := m.TextInput.Value()
				if jsonName == "" {
					jsonName = m.editingField.Name
				}
				m.editingField.JSONName = jsonName
				m.State = StateFieldEditMenu
				m.initializeFieldEditMenu()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateFieldEditMenu
				m.initializeFieldEditMenu()
			}
		}
	case StateFieldJSONName:
		m.TextInput, cmd = m.TextInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				jsonName := m.TextInput.Value()
				if jsonName == "" {
					jsonName = m.editingField.Name
				}
				m.editingField.JSONName = jsonName
				// Add field to the model
				m.editingModel.Fields = append(m.editingModel.Fields, *m.editingField)
				m.editingField = nil
				m.State = StateModelEdit
				m.initializeFieldMenu()
			} else if msg.Type == tea.KeyEsc {
				m.State = StateFieldType
				m.initializeFieldTypeList()
			}
		}

	case StateFieldList:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					for i, field := range m.editingModel.Fields {
						if field.Name == selectedItem.name {
							m.editingFieldIndex = i
							m.editingField = &m.editingModel.Fields[i]
							if m.removingField {
								// Remove the field
								m.editingModel.Fields = append(m.editingModel.Fields[:i], m.editingModel.Fields[i+1:]...)
								m.removingField = false
								m.State = StateModelEdit
								m.initializeFieldMenu()
							} else {
								// Edit the field
								m.State = StateFieldEditMenu
								m.initializeFieldEditMenu()
							}
							break
						}
					}
				}
			} else if msg.Type == tea.KeyEsc {
				m.State = StateModelEdit
				m.initializeFieldMenu()
			}
		}

	case StateModelList:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					if selectedItem.disabled {
						m.Err = fmt.Errorf("This option is not available.")
						return m, cmd
					}
					// Find the model index
					for i, model := range m.Data.Models {
						if model.Name == selectedItem.name {
							if m.removingModel {
								// Remove the model
								m.Data.Models = append(m.Data.Models[:i], m.Data.Models[i+1:]...)
								m.removingModel = false
								m.State = StateMainMenu
								m.initializeMainMenu()
							} else {
								// Edit the model
								m.editingModelIndex = i
								m.editingModel = &m.Data.Models[i]
								m.TextInput = textinput.New()
								m.TextInput.Placeholder = "Model Name"
								m.TextInput.SetValue(m.editingModel.Name)
								m.TextInput.Focus()
								m.State = StateModelEdit
								m.initializeFieldMenu()
							}
							break
						}
					}
				}
			} else if msg.Type == tea.KeyEsc {
				m.State = StateMainMenu
				m.initializeMainMenu()
			}
		}

	case StateDone:
		return m, tea.Quit
	}

	return m, cmd
} */

/* func (m *Model) initializeMainMenu() {
	items := []list.Item{
		common.Item{Name: "Add Service"},
		common.Item{Name: "Edit Service", Disabled: len(m.Data.Services) == 0},
		common.Item{Name: "Remove Service", Disabled: len(m.Data.Services) == 0},
		common.Item{Name: "Add Model"},
		common.Item{Name: "Edit Model", Disabled: len(m.Data.Models) == 0},
		common.Item{Name: "Remove Model", Disabled: len(m.Data.Models) == 0},
		common.Item{Name: "Add Repository", Disabled: true},
		common.Item{Name: "Edit Repository", Disabled: true},
		common.Item{Name: "Remove Repository", Disabled: true},
		common.Item{Name: "Add Handler", Disabled: true},
		common.Item{Name: "Edit Handler", Disabled: true},
		common.Item{Name: "Remove Handler", Disabled: true},
		common.Item{Name: "Save and Exit"},
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Blueprint Configuration"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.Err = nil // Reset any errors
}

// ################################################################################
// # Service Listing
// ################################################################################

func (m *Model) initializeServiceList() {
	items := []list.Item{}
	for _, svc := range m.Data.Services {
		items = append(items, common.Item{Name: svc.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Service"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.RemovingService = false // Reset removing flag
}

// ################################################################################
// # Service Editing Menu
// ################################################################################
func (m *Model) initializeServiceEditMenu() {
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
}

// ################################################################################
// # Service Type Listing
// ################################################################################
func (m *Model) initializeServiceTypeList() {
	items := []list.Item{
		common.Item{Name: "rest"},
		// Add other service types as needed
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select Service Type"
	m.List.SetSize(40, 15)
	m.List.Select(0)
}

// ################################################################################
// # Services's Model Listing
// ################################################################################
func (m *Model) initializeModelSelectionList() {
	items := []list.Item{}
	assignedModels := make(map[string]bool)
	for _, modelName := range m.EditingService.ModelNames {
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
}

// ################################################################################
// # Model Listing
// ################################################################################
func (m *Model) initializeModelList() {
	items := []list.Item{}
	for _, model := range m.Data.Models {
		items = append(items, common.Item{Name: model.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Model"
	m.List.SetSize(40, 15)
	m.List.Select(0)
	m.RemovingModel = false // Reset removing flag
}

// ################################################################################
// # Model's Field Menu
// ################################################################################
func (m *Model) initializeFieldMenu() {
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
}

// ################################################################################
// # Model's Field Listing
// ################################################################################
func (m *Model) initializeFieldList() {
	items := []list.Item{}
	for _, field := range m.EditingModel.Fields {
		items = append(items, common.Item{Name: field.Name})
	}
	m.List = list.New(items, common.CustomDelegate{}, 0, 0)
	m.List.Title = "Select a Field"
	m.List.SetSize(40, 15)
	m.List.Select(0)
}

// ################################################################################
// # Model's Field Editing Menu
// ################################################################################
func (m *Model) initializeFieldEditMenu() {
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
}

func (m *Model) initializeFieldTypeList() {
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
}

func getSelectedModelNames(items []list.Item) []string {
	var selectedModels []string
	for _, listItem := range items {
		if item, ok := listItem.(common.SelectableItem); ok && item.Selected {
			selectedModels = append(selectedModels, item.Name)
		}
	}
	return selectedModels
} */
