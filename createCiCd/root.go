package createCiCd

import (
	"brahma/data"
	"brahma/helper"
	"strings"
)

func CreateCiCdTemplate(config data.Config, args []string) {

	defer func() {
		if r := recover(); r != nil {
			panic("Not supported")
		}
	}()

	var template string
	var name string

	params := []string{config.CicdPipeline, args[0]}
	template = helper.FetchVariable(strings.Join(params, ""), variables).(string)

	if config.CicdPipeline == "jenkins" {
		name = "jenkinsfile"
	} else {
		name = "azure-pipeline.yaml"
	}

	helper.CreateFile(name, template)
}
