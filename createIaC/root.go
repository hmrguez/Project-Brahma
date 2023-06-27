package iac

import (
	"brahma/data"
	"brahma/helper"
	"fmt"
	"path"
)

func CreateIaCTemplate(config data.Config, args []string) {

	var template string
	if config.Infrastructure == "terraform" {

		if config.CloudProvider == "aws" {
			template = handleTerraformAws(config, args)
		} else if config.CloudProvider == "azure" {
			template = handleTerraformAzure(config, args)
		} else {
			panic("Not supported")
		}

	} else {
		panic("Not supported")
	}

	helper.CreateFile("template.tf", template)
}

func handleTerraformAws(config data.Config, args []string) string {
	var template string

	fmt.Println(args[0])

	if args[0] == "simple" {
		template = helper.ReadFile(path.Join("createIaC", "tf-aws-simple.tf"))
	} else {
		panic("Not supported")
	}

	return template
}

func handleTerraformAzure(config data.Config, args []string) string {
	var template string

	if args[0] == "simple" {
		template = helper.ReadFile(path.Join("createIaC", "tf-azure-simple.tf"))
	} else {
		panic("Not supported")
	}

	return template
}
