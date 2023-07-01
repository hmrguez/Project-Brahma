package createCont

import (
	"brahma/data"
	"brahma/helper"
	"strings"
)

func CreateContTemplate(config data.Config, args []string) {

	defer func() {
		if r := recover(); r != nil {
			panic("Not supported")
		}
	}()

	var template string

	params := []string{config.Containers, args[0]}
	template = helper.FetchVariable(strings.Join(params, ""), variables).(string)

	helper.CreateFile("dockerfile", template)
}
