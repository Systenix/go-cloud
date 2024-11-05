package cli

import (
	"fmt"
	"os"

	"github.com/Systenix/go-cloud/config"
	"github.com/Systenix/go-cloud/generators"
	tui_configure "github.com/Systenix/go-cloud/tui/configure_command"
	"github.com/Systenix/go-cloud/tui/configure_command/states"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var outputPath string

var configCmd = &cobra.Command{
	Use:   "configure",
	Short: "Create or edit a project configuration file that act as the blueprint for the project generator",
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

		var data *generators.ProjectData
		var err error

		if configPath != "" {
			parsedData, err := config.ParseConfig(configPath)
			if err != nil {
				fmt.Printf("Error parsing configuration file: %v\n", err)
				os.Exit(1)
			}
			// extract data from parsedData
			data = &generators.ProjectData{
				ProjectName: parsedData.ProjectName,
				ProjectPath: parsedData.ProjectPath,
				ModulePath:  parsedData.ModulePath,
				Services:    parsedData.Services,
				Models:      parsedData.Models,
				Events:      parsedData.Events,
			}
		} else {
			data = &generators.ProjectData{}
		}

		model := tui_configure.NewModel(data)
		model.SetState(states.NewProjectInfoState())
		p := tea.NewProgram(model)
		updatedModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}

		// access the updated data from the model
		updatedData := updatedModel.(*tui_configure.Model).Data

		// Save the configuration to the output file
		err = saveConfig(outputPath, updatedData)
		if err != nil {
			fmt.Printf("Error saving configuration file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Configuration saved to %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&configPath, "config-file", "c", "", "Path to the configuration file (YAML)")
	configCmd.Flags().StringVarP(&outputPath, "output", "o", "project_config.yaml", "Path to save the configuration file")
}

func saveConfig(filePath string, data generators.ProjectData) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, yamlData, 0644)
}
