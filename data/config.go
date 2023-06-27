package data

type Config struct {
	Infrastructure         string `json:"infrastructure"`
	Containers             string `json:"containers"`
	ContainerOrchestration string `json:"containerOrchestration"`
	CloudProvider          string `json:"cloudProvider"`
	CicdPipeline           string `json:"cicdPipeline"`
	ServerConfig           string `json:"serverConfig"`
	Monitoring             string `json:"monitoring"`
}