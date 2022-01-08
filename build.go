package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	cf "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

/*
 * Example Create Stack Code:
 * https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/cloudformation/CreateStack/CreateStack.go
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

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic(err)
	}

	// Create client for CLI
	cft := cf.NewFromConfig(cfg)

	// fmt.Println(cft.ListStacks(context.TODO(), &cf.ListStacksInput{}))
	listStacks(cft, cf.ListStacksInput{})
}
