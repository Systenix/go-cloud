package generate_command

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.State {

	case StateProjectName:
		m.TextInput, cmd = m.TextInput.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.Data.ProjectName = m.TextInput.Value()
				m.TextInput.Reset()
				m.TextInput.Placeholder = "Project Path (e.g., /github.com/Username)"
				m.State = StateProjectPath
			}
		}

	case StateProjectPath:
		m.TextInput, cmd = m.TextInput.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.Data.ProjectPath = m.TextInput.Value()
				m.TextInput.Blur()
				m.State = StateProtocol
				m.initializeProtocolList()
			}
		}

	case StateProtocol:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					m.Data.Protocol = selectedItem.name
					m.State = StateMessageBroker
					m.initializeMessageBrokerList()
				}
			}
		}

	case StateMessageBroker:
		var listCmd tea.Cmd
		m.List, listCmd = m.List.Update(msg)
		cmd = listCmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if selectedItem, ok := m.List.SelectedItem().(item); ok {
					m.Data.MessageBroker = selectedItem.name
					m.State = StateDone
					return m, tea.Quit
				}
			}
		}

	case StateDone:
		return m, tea.Quit
	}

	return m, cmd
}

type item struct {
	name string
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.name }

func (m *Model) initializeProtocolList() {
	items := []list.Item{
		item{name: "REST"},
		// Add more protocols if needed
	}
	m.List = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.List.Title = "Select a communication protocol:"
	m.List.SetSize(35, 15)
	m.List.Select(0)
}

func (m *Model) initializeMessageBrokerList() {
	items := []list.Item{
		item{name: "RabbitMQ"},
		item{name: "Kafka"},
		item{name: "NATS"},
	}
	m.List = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.List.Title = "Select a message broker:"
	m.List.SetSize(35, 15)
	m.List.Select(0)
}
