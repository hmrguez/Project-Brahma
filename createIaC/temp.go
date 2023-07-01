package createIaC

var variables = map[string]interface{}{
	"terraformazuresimple": terraformazuresimple,
	"terraformawssimple":   terraformawssimple,
}

var terraformazuresimple = `
	This is a multiline
	string for azure
`

var terraformawssimple = `
	This is a multiline
	string for aws
`
