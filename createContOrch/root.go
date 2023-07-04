package createContOrch

import (
	"brahma/data"
	"brahma/helper"
	"strings"
)

func CreateContOrchTemplate(config data.Config, args []string) {

	defer func() {
		if r := recover(); r != nil {
			panic("Not supported")
		}
	}()

	var template string
	var name = args[0] + ".yaml"

	params := []string{config.ContainerOrchestration, args[0]}
	template = helper.FetchVariable(strings.Join(params, ""), variables).(string)

	helper.CreateFile(name, template)
}
