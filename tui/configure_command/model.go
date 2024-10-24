package configure_command

import (
	"github.com/Systenix/go-cloud/generators"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	StateProjectInfo State = iota

	// Main Menu
	StateMainMenu

	// Service
	StateSelectServiceType
	StateServiceList
	// Service Edit
	StateServiceEdit
	StateServiceEditMenu
	StateEditServiceName
	// Service Model Assignment
	StateSelectModelsForService

	// Method
	StateMethodList
	StateMethodEdit

	// Handler
	StateHandlerList
	StateHandlerEdit

	// Model
	StateModelList
	StateModelEdit
	// Model's Field
	StateFieldList
	StateFieldEditMenu
	StateFieldName
	StateEditFieldName
	StateFieldType
	StateEditFieldType
	StateFieldJSONName
	StateEditFieldJSONName

	// Repository
	StateRepositoryList
	StateRepositoryEdit

	StateDone
)

// Model represents the state and data of the TUI config command.
type Model struct {
	// State represents the current state of the TUI application.
	State State

	// Data holds the project configuration data.
	Data generators.ProjectData

	// TextInput is the model for handling text input.
	TextInput textinput.Model

	// List is the model for handling list views.
	List list.Model

	// Cursor is used to track the current position in a list or menu.
	Cursor int

	// Err holds any error that occurs during the TUI execution.
	Err error

	// editingService is a pointer to the service currently being edited.
	editingService *generators.Service

	// editingServiceIndex is the index of the service currently being edited.
	editingServiceIndex int

	// removingService is a flag indicating whether a service is being removed.
	removingService bool

	// editingModel is a pointer to the model currently being edited.
	editingModel *generators.Model

	// editingField is a pointer to the field currently being edited.
	editingField *generators.Field

	// editingFieldIndex is the index of the field currently being edited.
	editingFieldIndex int

	// removingField is a flag indicating whether a field is being removed.
	removingField bool

	// editingModelIndex is the index of the model currently being edited.
	editingModelIndex int

	// removingModel is a flag indicating whether a model is being removed.
	removingModel bool

	// editingRepository is a pointer to the repository currently being edited.
	editingRepository *generators.Repository

	// editingHandler is a pointer to the handler currently being edited.
	editingHandler *generators.Handler

	// editingMethod is a pointer to the service method currently being edited.
	editingMethod *generators.ServiceMethod

	// Add other fields as needed
}

func InitialConfigModel(data *generators.ProjectData) Model {
	ti := textinput.New()
	ti.Placeholder = "Project Name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30

	return Model{
		State:     StateProjectInfo,
		TextInput: ti,
		Data:      *data,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}
