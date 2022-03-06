package main

import (
	"github.com/Fiddler25/cdk-go/lib/stacks"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := cdk.NewApp(nil)

	vpc, publicSubnet1 := stacks.CdkNetworkStack(app, "CdkNetworkStack", &cdk.StackProps{})
	stacks.CdkEc2Stack(app, "CdkEc2Stack", &stacks.CdkEc2StackProps{Vpc: vpc, PublicSubnet1: publicSubnet1})

	app.Synth(nil)
}
