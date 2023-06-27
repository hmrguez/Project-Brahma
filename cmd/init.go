package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	// Exported fields.
	Infrastructure         string `json:"infrastructure"`
	Containers             string `json:"containers"`
	ContainerOrchestration string `json:"containerOrchestration"`
	CloudProvider          string `json:"cloudProvider"`
	CicdPipeline           string `json:"cicdPipeline"`
	ServerConfig           string `json:"serverConfig"`
	Monitoring             string `json:"monitoring"`

	// Unexported fields.
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inits the brahma repository",
	Long:  `Inits the brahma repository with the default configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		myConfig := Config{
			Infrastructure:         "terraform",
			Containers:             "docker",
			ContainerOrchestration: "kubernetes",
			CloudProvider:          "aws",
			CicdPipeline:           "jenkins",
			ServerConfig:           "ansible",
			Monitoring:             "grafana",
		}

		/* TODO: Here will go any flag to change config */

		jsonConfig, _ := json.MarshalIndent(myConfig, "", "    ")

		file, err := os.Create("brahma.config")

		if err != nil {
			fmt.Println("There was an error creating the config file")
		}

		defer file.Close()

		err = ioutil.WriteFile("brahma.config", jsonConfig, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Brahma repo created successfully.")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
