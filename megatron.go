package main

import "text/template"

type Megatron struct {
	cloudFormationTemplate string
}

func getCloudFormationTemplate(templatePath string) (*template.Template, error) {
	t, err := template.ParseFiles(templatePath)

	return t, err
}

func NewMegatron(templatePath, userDataPath string) (Megatron, error) {
	m := Megatron{
		cloudFormationTemplate: templatePath,
	}

	return m, nil
}
