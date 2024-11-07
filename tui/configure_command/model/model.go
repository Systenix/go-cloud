package model

import (
	"github.com/Systenix/go-cloud/generators"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// State interface defines the Init, Update and View methods that each state must implement.
type State interface {
	Init(m *Model) tea.Cmd
	Update(msg tea.Msg, m *Model) tea.Cmd
	View(m *Model) string
}

// Model represents the state and data of the TUI config command.
type Model struct {
	// State represents the current state of the TUI application.
	CurrentState State

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

	// EditingService is a pointer to the service currently being edited.
	EditingService *generators.Service

	// EditingServiceIndex is the index of the service currently being edited.
	EditingServiceIndex int

	// RemovingService is a flag indicating whether a service is being removed.
	RemovingService bool

	// EditingModel is a pointer to the model currently being edited.
	EditingModel *generators.Model

	// EditingField is a pointer to the field currently being edited.
	EditingField *generators.Field

	// EditingFieldIndex is the index of the field currently being edited.
	EditingFieldIndex int

	// RemovingField is a flag indicating whether a field is being removed.
	RemovingField bool

	// EditingModelIndex is the index of the model currently being edited.
	EditingModelIndex int

	// RemovingModel is a flag indicating whether a model is being removed.
	RemovingModel bool

	// EditingRepository is a pointer to the repository currently being edited.
	EditingRepository *generators.Repository

	// EditingHandler is a pointer to the handler currently being edited.
	EditingHandler *generators.Handler

	// EditingMethod is a pointer to the service method currently being edited.
	EditingMethod *generators.ServiceMethod

	// Add other fields as needed
}

// SetState transitions the model to a new state.
/* func (m *Model) SetState(state State) {
	m.CurrentState = state
	cmd := state.Init(m)
	if cmd != nil {
		tea.Batch(cmd)
	}
}
*/

func (m *Model) SetState(state State) tea.Cmd {
	m.CurrentState = state
	return state.Init(m)
}
