package iac

var variables = map[string]interface{}{
	"tfazuresimple": tfazuresimple,
	"tfawssimple":   tfawssimple,
}

var tfazuresimple = `
	This is a multiline
	string for azure
`

var tfawssimple = `
	This is a multiline
	string for aws
`
