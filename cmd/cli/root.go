package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cloud-cli",
	Short: "go-cloud is a CLI tool for generating microservices in Go",
	Long: `
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

go-cloud-cli automates the creation of microservices.
It reads and create configuration files to generate domain models, handlers, and scaffolding code.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
