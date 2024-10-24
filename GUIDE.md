# Comprehensive Guide to go-cloud

## Introduction

Welcome to the comprehensive guide for **go-cloud**, a tool designed to simplify microservices development using Go and Clean Architecture principles. This guide will walk you through everything you need to know to make the most out of go-cloud.

## Prerequisites

- **Go Programming Language**: Version 1.16 or higher.
- **Basic Knowledge of Go**: Understanding of Go modules and packages.
- **Familiarity with Microservices**: Concepts and architecture.

## Installation

### **Step 1: Install Go**

If you haven't already, download and install Go from the [official website](https://golang.org/dl/).

### **Step 2: Install go-cloud**

```bash
go install github.com/Systenix/go-cloud/cmd/go-cloud@latest
```

Ensure your $GOPATH/bin is in your system’s PATH to run go-cloud from any directory.

## Getting Started

### Configuration Command (config)

The config command launches an interactive TUI to configure your project.

#### Navigating the TUI

	•	Main Menu: Access options to add/edit services, models, repositories, etc.
	•	Controls:
	•	Enter: Select or confirm.
	•	Esc: Go back or cancel.
	•	Ctrl+C: Exit the TUI.

#### Configuring Services

	•	Add Service: Define a new service.
	•	Edit Service: Modify existing services.
	•	Assign Models: Link models to services.

#### Configuring Models

	•	Add Model: Create a new data model.
	•	Edit Model: Modify existing models.
	•	Add Fields: Define fields with names, types, and JSON names.
	•	Nested Fields: Use existing models as field types for nesting.

### Generation Command (generate)

After configuring your project, generate the code:

```bash
go-cloud generate -c <path/to/config.yaml>
```

This command reads the configuration and generates the project structure in the current directory.

### Understanding the Generated Structure

	•	cmd/: Entry points for your services.
	•	internal/: Core application code.
	•	go.mod/: Go module file for the project.

### Advanced Usage

#### Coming soon

- Ability to customize the code generation templates.
- Adding Complex Models
	-	Use arrays and maps in field types.
	-	Implement relationships between models.

### Best Practices

	•	Consistent Naming: Use clear and descriptive names for services and models.
	•	Modularity: Keep services focused and modular.
	•	Version Control: Commit your configuration files to version control.

### Troubleshooting

#### Common Issues

	•	Command Not Found: Ensure go-cloud is in your PATH.
	•	Permission Denied: Check file permissions and run with appropriate privileges.
	•	Unexpected Errors: Run with the --debug flag for more information.

#### FAQ

Q: Can I use go-cloud with existing projects?

A: Yes, you can integrate go-cloud into existing projects by carefully merging the generated code.

Q: Does go-cloud support gRPC or GraphQL?

A: Currently, go-cloud supports REST services. Support for additional protocols is on the roadmap.

### Conclusion

With go-cloud, you can accelerate your microservices development while adhering to clean architecture principles. We hope this guide has been helpful. Happy coding!