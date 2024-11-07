package states

import (
	"github.com/Systenix/go-cloud/tui/configure_command/common"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectInfoState struct{}

func NewProjectInfoState() *ProjectInfoState {
	return &ProjectInfoState{}
}

func (s *ProjectInfoState) Init(m *model.Model) tea.Cmd {
	m.TextInput = textinput.New()
	m.TextInput.Placeholder = "Project Name"
	m.TextInput.Focus()
	m.TextInput.CharLimit = 64
	m.TextInput.Width = 30
	return textinput.Blink
}

func (s *ProjectInfoState) Update(msg tea.Msg, m *model.Model) tea.Cmd {
	var cmd tea.Cmd
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
				m.TextInput.Reset()
				m.SetState(NewMainMenuState())
			}
		}
	}
	return cmd
}

func (s *ProjectInfoState) View(m *model.Model) string {
	if m.Data.ProjectName == "" {
		return common.RenderPrompt("Enter the project name:", m.TextInput.View())
	} else if m.Data.ModulePath == "" {
		return common.RenderPrompt("Enter the module path (e.g., github.com/Username/project):", m.TextInput.View())
	}
	return ""
}
