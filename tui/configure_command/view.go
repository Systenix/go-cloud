package configure_command

import (
	"fmt"
	"strings"

	"github.com/Systenix/go-cloud/tui/configure_command/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (m *Model) View() string {
	var mainContent string

	if m.CurrentState != nil {
		mainContent = m.CurrentState.View(m)
	}

	statusBar := lipgloss.NewStyle().
		Background(lipgloss.Color("#626262")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1).
		Render(fmt.Sprintf("State: %T", m.CurrentState))

	// Render the blueprint state
	blueprint := m.renderBlueprint()

	// Create side-by-side layout
	mainColumn := styles.ColumnStyle.Width(50).Render(mainContent)
	blueprintColumn := styles.ColumnStyle.Width(30).Render(blueprint)
	content := lipgloss.JoinHorizontal(lipgloss.Top, mainColumn, styles.Divider.String(), blueprintColumn)

	return lipgloss.JoinVertical(lipgloss.Top, content, statusBar)
}

func (m *Model) renderBlueprint() string {
	var sb strings.Builder

	sb.WriteString(styles.TitleStyle.Render("Blueprint State") + "\n\n")
	sb.WriteString(fmt.Sprintf("%s: %s\n", lipgloss.NewStyle().Bold(true).Render("Project Name"), m.Data.ProjectName))
	sb.WriteString(fmt.Sprintf("%s: %s\n", lipgloss.NewStyle().Bold(true).Render("Module Path"), m.Data.ModulePath))
	sb.WriteString("\n")
	sb.WriteString(styles.TitleStyle.Render("Services:") + "\n")

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
