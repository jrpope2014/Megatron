package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/config"
	cf "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

/*
 * Example List Stack Code:
 * https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/cloudformation/ListStacks/ListStacks.go
 */
func listStacks(cft *cf.Client, lsi cf.ListStacksInput) {
	resp, err := cft.ListStacks(context.TODO(), &lsi)

	if err != nil {
		log.Panic("Something went wrong when trying to list cf stacks!")
	}

	fmt.Println("------Stack Summaries------")
	for _, stack := range resp.StackSummaries {
		fmt.Printf("Stack Name:%s\nStack Status:%s\nStack Creation Time:%s\n\n",
			*stack.StackName, stack.StackStatus, stack.CreationTime)
	}

	fmt.Println("-----------")
}

/*
 * Example Create Stack Code:
 * https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/cloudformation/CreateStack/CreateStack.go
 */
func createStack(cft *cf.Client, templateBody, stackName *string) *cf.CreateStackOutput {
	resp, err := cft.CreateStack(context.TODO(), &cf.CreateStackInput{
		TemplateBody: templateBody,
		StackName:    stackName,
	})

	if err != nil {
		log.Fatalf("CreateStack failed with error: %s", err)
	}

	return resp
}

/*
 * Example Delete Stack Code:
 * https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/cloudformation/DeleteStack/DeleteStack.go
 */
func deleteStack(cft *cf.Client, stackName *string) *cf.DeleteStackOutput {
	resp, err := cft.DeleteStack(
		context.TODO(),
		&cf.DeleteStackInput{StackName: stackName})

	if err != nil {
		log.Fatalf("DeleteStack failed with error: %s", err)
	}

	return resp
}

type UserData struct {
	Script      string
	ScriptLines []string
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic(err)
	}

	var command string
	flag.StringVar(&command, "c", "LIST", "Megatron command to use")
	flag.Parse()

	fmt.Printf("Command has a value of: %s\n", command)

	// Create client for CLI
	cft := cf.NewFromConfig(cfg)

	if command == "LIST" {
		// fmt.Println(cft.ListStacks(context.TODO(), &cf.ListStacksInput{}))
		listStacks(cft, cf.ListStacksInput{})
	}

	if command == "CREATE" {
		stackName := "MegatronTestStack"

		template, err := ioutil.ReadFile("cf_templates/testStack.yaml")

		if err != nil {
			log.Fatalf("Failed to read template file with error: %s", err)
		}

		templateBody := string(template)

		resp := createStack(cft, &templateBody, &stackName)

		log.Printf("Stack Creation Respose: %s", resp.ResultMetadata)
	}

	if command == "DELETE" {
		stackName := "MegatronTemplateTestStack"

		resp := deleteStack(cft, &stackName)

		log.Printf("Stack Deletion Response: %s", resp.ResultMetadata)
	}

	if command == "TEMPLATE_TEST" {
		templatePath := "cf_templates/templateTest.yaml"
		t, err := template.ParseFiles(templatePath)
		var body bytes.Buffer

		if err != nil {
			log.Fatalf("Failed to read %s with error %s", templatePath, err)
		}

		data, err := os.ReadFile("user-data.sh")

		encodedData := base64.StdEncoding.EncodeToString(data)

		if err != nil {
			log.Fatalf("%s", err)
		}

		testScript := UserData{
			Script:      encodedData,
			ScriptLines: strings.Split(string(encodedData), "\n"),
		}

		err = t.Execute(&body, testScript)

		if err != nil {
			log.Fatalf("%s", err)
		}

		templateBody := body.String()

		fmt.Println(testScript.ScriptLines)
		fmt.Println(templateBody)
		templateName := "MegatronTemplateTestStack"
		createStack(cft, &templateBody, &templateName)
	}
}
