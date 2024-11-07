package configure_command

import (
	"github.com/Systenix/go-cloud/generators"
	"github.com/Systenix/go-cloud/tui/configure_command/model"
	tea "github.com/charmbracelet/bubbletea"
)

type ConfigureModel struct {
	*model.Model
}

func (m *ConfigureModel) Init() tea.Cmd {
	if m.CurrentState != nil {
		return m.CurrentState.Init(m.Model)
	}
	return nil
}

// NewModel creates a new model with the provided project data.
func NewModel(data *generators.ProjectData) *ConfigureModel {
	return &ConfigureModel{
		Model: &model.Model{
			Data: *data,
		},
	}
}
