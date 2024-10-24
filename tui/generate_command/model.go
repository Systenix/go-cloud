package generate_command

import (
	"github.com/Systenix/go-cloud/generators"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	StateProjectName State = iota
	StateProjectPath
	StateProtocol
	StateMessageBroker
	StateDone
)

type Model struct {
	State     State
	Data      generators.ProjectData
	TextInput textinput.Model
	List      list.Model
	Cursor    int
	Err       error
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Project Name"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 101

	return Model{
		State:     StateProjectName,
		TextInput: ti,
		Data:      generators.ProjectData{},
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
