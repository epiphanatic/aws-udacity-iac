package main

// reference: https://www.go-on-aws.com/aws-go-sdk-v2/sdkv2/parms/

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var client *cloudformation.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = cloudformation.NewFromConfig(cfg)

	// params := &ec2.DescribeInstancesInput{
	// 	MaxResults: aws.Int32(10),
	// }

	params := &cloudformation.CreateStackInput{
		StackName:    aws.String("udagram-network-stack"),
		TemplateBody: aws.String("file://networking-and-infrastructure/network.yml"),
		Parameters:   aws.String("file://networking-and-infrastructure/network-params.json"),
	}

	client.CreateStack(context.TODO())
}

func main() {

	// cloudformation.CreateStack(stackInput)

	// Println prints with a new line
	fmt.Println("Hello, World!")

	const x int = 5
	fmt.Println(x)
	fmt.Println("hey", x, "there")
	// v is a variable value. T is type. See more here: https://pkg.go.dev/fmt
	// note that this is what you use if you dont want auto spacing
	fmt.Printf("hey%vthere\n", x)

	// ask user for their name
	var name string
	// need to use & to get the address of the variable (pointer) when using
	//  fmt.Scan to get user input
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)

	// print out the name
	fmt.Println("Hello", name)
}
