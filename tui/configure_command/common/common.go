package common

import (
	"fmt"
	"io"

	"github.com/Systenix/go-cloud/tui/configure_command/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Item represents an item in the list.
type Item struct {
	Name     string
	Disabled bool
}

func (i Item) Title() string {
	if i.Disabled {
		return styles.DisabledItemStyle.Render(i.Name)
	}
	return i.Name
}

func (i Item) Description() string { return "" }
func (i Item) FilterValue() string { return i.Name }

type SelectableItem struct {
	Name     string
	Selected bool
}

func (i SelectableItem) Title() string {
	if i.Selected {
		return "[x] " + i.Name
	}
	return "[ ] " + i.Name
}

func (i SelectableItem) Description() string { return "" }
func (i SelectableItem) FilterValue() string { return i.Name }

// CustomDelegate implements the list item delegate.
type CustomDelegate struct{}

func (d CustomDelegate) Height() int                               { return 1 }
func (d CustomDelegate) Spacing() int                              { return 0 }
func (d CustomDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d CustomDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(Item)
	if !ok {
		return
	}

	var style lipgloss.Style
	if index == m.Index() {
		if item.Disabled {
			style = styles.DisabledItemStyle.Copy().Bold(true)
		} else {
			style = styles.SelectedItemStyle
		}
	} else {
		if item.Disabled {
			style = styles.DisabledItemStyle
		} else {
			style = lipgloss.NewStyle()
		}
	}

	fmt.Fprintf(w, style.Render(item.Title()))
}

// RenderPrompt renders a prompt with a text input and a continue button.
func RenderPrompt(prompt, inputView string) string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		styles.PromptStyle.Render(prompt),
		inputView,
		"(Press Enter to continue)",
	)
}
