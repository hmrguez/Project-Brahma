package iac

import (
	"brahma/data"
	"brahma/helper"
	"strings"
)

func CreateIaCTemplate(config data.Config, args []string) {

	var template string
	if config.Infrastructure == "terraform" {
		params := []string{"tf", config.CloudProvider, args[0]}
		template = helper.FetchVariable(strings.Join(params, ""), variables).(string)
	} else {
		panic("Not supported")
	}

	helper.CreateFile("template.tf", template)
}
