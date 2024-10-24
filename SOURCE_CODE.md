# Source Code Overview

## Introduction

This document provides an in-depth look at the source code of **go-cloud**, explaining the architecture and design choices made during development. Whether you're a contributor or just curious about the inner workings, this guide will help you understand how the project is structured and how it functions.

## Project Structure

The go-cloud project is organized into several packages and directories:

- **`cmd/go-cloud`**: Contains the main executable code for the CLI.
- **`generators`**: Houses the code generation logic.
- **`config_command`**: Implements the configuration command and the TUI.
- **`templates`**: Contains templates used for code generation.
- **`models`**: Defines the data structures used throughout the project.

```bash
go-cloud
├── GUIDE.md
├── LICENSE
├── Makefile
├── README.md
├── SOURCE_CODE.md
├── bin
│   └── go-cloud-cli
├── cmd
│   ├── cli
│   │   ├── config.go
│   │   ├── generate.go
│   │   └── root.go
│   └── main.go
├── config
│   └── config_parser.go
├── docs
├── examples
│   └── game_service_blueprint.yaml
├── generators
│   └── project_generator.go
├── go.mod
├── go.sum
├── scripts
├── templates
│   ├── Dockerfile.tmpl
│   ├── cmd
│   │   └── main.go.tmpl
│   ├── docker-compose.yml.tmpl
│   ├── embed.go
│   ├── events
│   │   └── event.go.tmpl
│   ├── go.mod.tmpl
│   ├── handlers
│   │   └── handler.go.tmpl
│   ├── models
│   │   └── model.go.tmpl
│   ├── repositories
│   │   └── repository.go.tmpl
│   └── services
│       └── service.go.tmpl
└── tui
    ├── config_command
    │   ├── model.go
    │   ├── styles.go
    │   ├── update.go
    │   └── view.go
    └── generate_command
        ├── model.go
        ├── update.go
        └── view.go

19 directories, 32 files
```

## CLI Implementation with Cobra

The CLI is built using [Cobra](https://github.com/spf13/cobra), a library for creating powerful command-line applications in Go.

### **Key Commands**

- **`config`**: Launches the interactive TUI for configuration.
- **`generate`**: Generates the project scaffolding based on the configuration.

### **Command Structure**

Each command is defined in its own file within the `cmd/cli` directory.

```go
// cmd/main.go
package main

import "github.com/Systenix/go-cloud/cmd/cli"

func main() {
	cli.Execute()
}
```

```go
// cmd/cli/generate.go
package cli

import (
	"fmt"
	"os"

	"github.com/Systenix/go-cloud/config"
	"github.com/Systenix/go-cloud/generators"
	tui_generate "github.com/Systenix/go-cloud/tui/generate_command"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var configPath string

var generateCmd = &cobra.Command{...}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file (YAML/JSON)")
}

```

In the _init()_ function, the generate command is added to the root command and the config flag assigned to the generate command

## State Machine and TUI Implementation

### **State Machine Pattern**

The TUI uses a state machine to manage different views and user interactions. This pattern allows for a clear separation of concerns and easier maintenance.

    •	States: Defined constants representing each state of the TUI (e.g., StateMainMenu, StateServiceEdit).
    •	Model: Holds the current state and data.
    •	Update Function: Contains logic to transition between states based on user input.

#### Implementation with Bubble Tea

[Bubble Tea](https://github.com/charmbracelet/bubbletea) is used to build the TUI. It provides a functional and reactive framework inspired by The Elm Architecture.

#### Key Components

	•	Model: Represents the state of the application.
	•	Msg: Represents messages/events that trigger state changes.
	•	Update Function: Updates the model based on messages.
	•	View Function: Renders the UI based on the current state.

### Styling with Lip Gloss

[Lip Gloss](https://github.com/charmbracelet/lipgloss) is used for styling the TUI components.

	•	Styles: Define the look and feel of UI elements.
	•	Layout: Arrange components using Lip Gloss’s layout utilities.

### Key Components and Modules

*tui/config_command package*

Implements the configuration command and TUI logic.

	•	Model Definitions: Holds the state and data for the TUI.
	•	State Management: Update and View functions handle state transitions and rendering.

*tui/generate_command package*

Implements the generate command and TUI logic. Follow the same logic as the config_command.

*generators package*

Handles code generation based on the configuration.

	•	Template Parsing: Uses text/template to process templates.
	•	File Generation: Writes generated code to the appropriate directories.

Defines data structures used across the application, such as:

	•	Service
	•	Model
	•	Field
	•	Repository
	•	Handler

## Conclusion

This overview should give you a solid understanding of the go-cloud source code. If you have any questions or need further clarification, feel free to open an issue or reach out directly.