package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic(err)
	}

	var command string
	flag.StringVar(&command, "c", "LIST", "Megatron command to use")
	flag.Parse()

	fmt.Printf("Command has a value of: %s\n", command)

	defaultTemplatePath := "cf_templates/testStack.yaml"
	defaultUserDataPath := "user-data.sh"

	// Create client for CLI
	cft := cf.NewFromConfig(cfg)

	if command == "LIST" {
		// fmt.Println(cft.ListStacks(context.TODO(), &cf.ListStacksInput{}))
		listStacks(cft, cf.ListStacksInput{})
	}

	if command == "CREATE" {
		stackName := "MegatronTestStack"

		m := NewMegatron(defaultTemplatePath, defaultUserDataPath)

		resp := createStack(cft, &m.CloudFormationTemplate, &stackName)

		log.Printf("Stack Creation Respose: %s", resp.ResultMetadata)
	}

	if command == "DELETE" {
		stackName := "MegatronTemplateTestStack"

		resp := deleteStack(cft, &stackName)

		log.Printf("Stack Deletion Response: %s", resp.ResultMetadata)
	}

}
