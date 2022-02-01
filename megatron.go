package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"os"
	"text/template"
)

type Megatron struct {
	CloudFormationTemplate string
	UserData               string
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

func NewMegatron(templatePath, userDataPath string) Megatron {

	t, err := getCloudFormationTemplate(templatePath)

	if err != nil {
		log.Fatal(err)
	}

	userData, err := loadEncodedUserData(userDataPath)

	m := Megatron{
		CloudFormationTemplate: templatePath,
		UserData:               userData,
	}

	if err != nil {
		log.Fatal(err)
	}

	var cft bytes.Buffer

	err = t.Execute(&cft, m)

	if err != nil {
		log.Fatal(err)
	}

	m.CloudFormationTemplate = cft.String()

	return m
}
