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

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new microservice project based on the provided configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
                                           ,--,
                                         ,--.'|                               ,---,
                ,---.      ,---,.        |  | :     ,---.           ,--,    ,---.'|
    ,----._,.  '   ,'\   ,'  .' |        :  : '    '   ,'\        ,'_ /|    |   | :
   /   /  ' / /   /   |,---.'   , ,---.  |  ' |   /   /   |  .--. |  | :    |   | |
  |   :     |.   ; ,. :|   |    |/     \ '  | |  .   ; ,. :,'_ /| :  . |  ,--.__| |
  |   | .\  .'   | |: ::   :  .'/    / ' |  | :  '   | |: :|  ' | |  . . /   ,'   |
  .   ; ';  |'   | .; ::   |.' .    ' /  '  : |__'   | .; :|  | ' |  | |.   '  /  |
  '   .   . ||   :    |'---'   '   ; :__ |  | '.'|   :    |:  | : ;  ; |'   ; |:  |
   '---'-'| | \   \  /         '   | '.'|;  :    ;\   \  / '  :  '--'   \   | '/  '
   .'__/\_: |  '----'          |   :    :|  ,   /  '----'  :  ,      .-./   :    :|
   |   :    :                   \   \  /  ---'-'            '--'----'    \   \  /
    \   \  /                     '----'                                   '----'
     '--'-'
    `)
		// Start the TUI to collect basic project info
		p := tea.NewProgram(tui_generate.InitialModel())

		model, err := p.Run()
		if err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}

		// Access the collected data
		data := model.(tui_generate.Model).Data

		// Parse the configuration file if provided
		if configPath != "" {
			configData, err := config.ParseConfig(configPath)
			if err != nil {
				fmt.Printf("Error parsing configuration file: %v\n", err)
				os.Exit(1)
			}
			// Map configData to ProjectData
			data.Services = configData.Services
			data.Models = configData.Models
			data.Events = configData.Events

		} else {
			fmt.Println("No configuration file provided. Please provide a valid blueprint config file.")
			os.Exit(1)
		}

		// Generate the project with the collected data
		err = generators.GenerateProject(data)
		if err != nil {
			fmt.Println("Error generating project:", err)
			os.Exit(1)
		}

		fmt.Printf("Project %s has been successfully created at %s.\n", data.ProjectName, data.ProjectPath)
	},
}

func init() {
	// Add the generate command to the root command
	rootCmd.AddCommand(generateCmd)
	// Add the config flag to the generate command
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file (YAML/JSON)")
}
