package configure_command

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (m Model) View() string {
	var mainContent string

	switch m.State {
	case StateProjectInfo:
		if m.Data.ProjectName == "" {
			mainContent = renderPrompt("Enter the project name:", m.TextInput.View())
		} else if m.Data.ModulePath == "" {
			mainContent = renderPrompt("Enter the module path (e.g., github.com/Username/project):", m.TextInput.View())
		}

	case StateMainMenu:
		if m.Err != nil {
			errorMsg := errorStyle.Render(m.Err.Error())
			mainContent = fmt.Sprintf("%s\n\n%s", m.List.View(), errorMsg)
			m.Err = nil // Reset the error after displaying
		} else {
			mainContent = m.List.View()
		}

	case StateServiceEdit:
		if m.editingService.Name == "" {
			mainContent = renderPrompt("Enter the service name:", m.TextInput.View())
		} else if m.editingService.Type == "" {
			mainContent = renderPrompt("Enter the service type (e.g., rest):", m.TextInput.View())
		} else {
			mainContent = m.List.View()
		}

	case StateServiceEditMenu:
		mainContent = m.List.View()

	case StateEditServiceName:
		helpText := "\n\nPress Enter to continue, Esc to cancel."
		mainContent = renderPrompt("Enter the new service name:", m.TextInput.View()) + helpText

	case StateSelectServiceType:
		mainContent = m.List.View()

	case StateSelectModelsForService:
		helpText := "\n\nPress Space to toggle selection, Enter to confirm, Esc to go back."
		mainContent = m.List.View() + helpText

	case StateServiceList:
		mainContent = m.List.View()

	case StateModelEdit:
		if m.editingModel.Name == "" {
			mainContent = renderPrompt("Enter the model name:", m.TextInput.View())
		} else {
			mainContent = m.List.View()
		}

	case StateFieldList:
		mainContent = m.List.View()

	case StateFieldEditMenu:
		mainContent = m.List.View()

	case StateEditFieldName:
		mainContent = renderPrompt("Enter the new field name:", m.TextInput.View())

	case StateFieldName:
		mainContent = renderPrompt("Enter the field name:", m.TextInput.View())

	case StateEditFieldType:
		mainContent = m.List.View()

	case StateFieldType:
		mainContent = m.List.View()

	case StateEditFieldJSONName:
		mainContent = renderPrompt("Enter the JSON name (leave empty to use field name):", m.TextInput.View())

	case StateFieldJSONName:
		mainContent = renderPrompt("Enter the JSON name (leave empty to use field name):", m.TextInput.View())

	case StateModelList:
		mainContent = m.List.View()

	case StateDone:
		mainContent = "Configuration complete. Saving..."

	default:
		mainContent = "Unknown state"
	}

	statusBar := lipgloss.NewStyle().
		Background(lipgloss.Color("#626262")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1).
		Render(fmt.Sprintf("State: %d", m.State)) // Here, we want to have a progress bar

	// Render the blueprint state
	blueprint := m.renderBlueprint()

	// Create side-by-side layout
	mainColumn := columnStyle.Width(50).Render(mainContent)
	blueprintColumn := columnStyle.Width(30).Render(blueprint)
	content := lipgloss.JoinHorizontal(lipgloss.Top, mainColumn, divider.String(), blueprintColumn)

	return lipgloss.JoinVertical(lipgloss.Top, content, statusBar)
}

func (m *Model) renderBlueprint() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Blueprint State") + "\n\n")
	sb.WriteString(fmt.Sprintf("%s: %s\n", lipgloss.NewStyle().Bold(true).Render("Project Name"), m.Data.ProjectName))
	sb.WriteString(fmt.Sprintf("%s: %s\n", lipgloss.NewStyle().Bold(true).Render("Module Path"), m.Data.ModulePath))
	sb.WriteString("\n")
	sb.WriteString(lipgloss.NewStyle().Bold(true).Render("Services:") + "\n")

	for _, svc := range m.Data.Services {
		sb.WriteString(fmt.Sprintf("  • %s (%s)\n", svc.Name, svc.Type))
		if len(svc.ModelNames) > 0 {
			sb.WriteString("    Models:\n")
			for _, modelName := range svc.ModelNames {
				sb.WriteString(fmt.Sprintf("      - %s\n", modelName))
			}
		}
	}

	sb.WriteString("\n")
	sb.WriteString(lipgloss.NewStyle().Bold(true).Render("Models:") + "\n")
	for _, model := range m.Data.Models {
		sb.WriteString(fmt.Sprintf("  • %s\n", model.Name))
		for _, field := range model.Fields {
			sb.WriteString(fmt.Sprintf("    - %s: %s\n", field.Name, field.Type))
		}
	}
	// Wrap text to fit the column width
	return wordwrap.String(sb.String(), 30)
}

// renderPrompt renders a prompt with a text input and a continue button.
func renderPrompt(prompt, inputView string) string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		promptStyle.Render(prompt),
		inputView,
		"(Press Enter to continue)",
	)
}

type customDelegate struct{}

func (d customDelegate) Height() int                               { return 1 }
func (d customDelegate) Spacing() int                              { return 0 }
func (d customDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d customDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(item)
	if !ok {
		return
	}

	var style lipgloss.Style
	if index == m.Index() {
		if item.disabled {
			style = disabledItemStyle.Copy().Bold(true)
		} else {
			style = selectedItemStyle
		}
	} else {
		if item.disabled {
			style = disabledItemStyle
		} else {
			style = lipgloss.NewStyle()
		}
	}

	fmt.Fprintf(w, style.Render(item.Title()))
}
