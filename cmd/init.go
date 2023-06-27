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

var inf string
var cont string
var contOrch string
var cProv string
var cicd string
var servConf string
var monit string

var initCmd = &cobra.Command{

	Use:   "init",
	Short: "Inits the brahma repository",
	Long:  `Inits the brahma repository with the default configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		myConfig := Config{
			Infrastructure: inf,
			Containers: cont,
			ContainerOrchestration: contOrch,
			CloudProvider: cProv,
			CicdPipeline: cicd,
			ServerConfig: servConf,     
			Monitoring: monit,
		}

		jsonConfig, _ := json.MarshalIndent(myConfig, "", "    ")

		file, err := os.Create("brahma.config")

		if err != nil {
			fmt.Println("There was an error creating the config file: ", err)
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
	initCmd.Flags().StringVarP(&inf, "iac", "i", "terraform", "IaC tool to use")
	initCmd.Flags().StringVarP(&cont, "cont", "c", "docker", "Containerization tool to use")
	initCmd.Flags().StringVarP(&contOrch, "cont-orch", "o", "kubernetes", "Kubernetes tool to use")
	initCmd.Flags().StringVarP(&cProv, "cloud-prov", "p", "aws", "Cloud Provider tool to use")
	initCmd.Flags().StringVarP(&monit, "monitor", "m", "grafana", "Monitoring tool to use")
	initCmd.Flags().StringVarP(&cicd, "cicd", "d", "jenkins", "CI/CD pipeline tool to use")
	initCmd.Flags().StringVarP(&servConf, "serv-config", "s", "ansible", "Server configuration tool to use")
	rootCmd.AddCommand(initCmd)
}
