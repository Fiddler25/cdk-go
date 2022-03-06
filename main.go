package main

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewCdkGoStack(scope constructs.Construct, id string, props *cdk.StackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = *props
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	awssqs.NewQueue(stack, jsii.String("CdkGoQueue"), &awssqs.QueueProps{
		VisibilityTimeout: cdk.Duration_Seconds(jsii.Number(300)),
	})

	return stack
}

func main() {
	app := cdk.NewApp(nil)

	NewCdkGoStack(app, "CdkGoStack", &cdk.StackProps{})

	app.Synth(nil)
}
