package iac

import (
	"brahma/data"
	"brahma/helper"
	"strings"
)

func CreateIaCTemplate(config data.Config, args []string) {

	defer func() {
		if r := recover(); r != nil {
			panic("Not supported")
		}
	}()

	var template string
	
	params := []string{config.Infrastructure, config.CloudProvider, args[0]}
	template = helper.FetchVariable(strings.Join(params, ""), variables).(string)

	helper.CreateFile("template.tf", template)
}
