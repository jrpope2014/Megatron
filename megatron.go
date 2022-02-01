package main

import (
	"encoding/base64"
	"log"
	"os"
	"text/template"
)

type Megatron struct {
	cloudFormationTemplate string
}

func getCloudFormationTemplate(templatePath string) (*template.Template, error) {
	t, err := template.ParseFiles(templatePath)

	return t, err
}

func loadEncodedUserData(userDataPath string) (string, error) {
	data, err := os.ReadFile(userDataPath)
	if err != nil {
		panic(err)
	}

	encodedData := base64.StdEncoding.EncodeToString(data)

	return encodedData, err
}

func NewMegatron(templatePath, userDataPath string) (Megatron, error) {

	t, err := getCloudFormationTemplate(templatePath)

	if err != nil {
		log.Fatal(err)
	}

	m := Megatron{
		cloudFormationTemplate: templatePath,
	}

	return m, nil
}
