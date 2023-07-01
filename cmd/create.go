package cmd

import (
	"brahma/createCont"
	"brahma/createIaC"
	"brahma/data"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create templates",
	Long:  `Create templates for any configuration type`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("No arguments passed to create command")
			return
		}

		// Open the file for reading
		file, err := os.Open("brahma.config")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Decode the JSON data
		var config data.Config
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)

		switch args[0] {
		case "iac":
			createIaC.CreateIaCTemplate(config, args[1:])
		case "cont":
			createCont.CreateContTemplate(config, args[1:])
		case "cont-orch":
			panic("No implementation")
		default:
			panic("No implementation")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
