package main

import (
	"github.com/Fiddler25/cdk-go/lib/stacks"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := cdk.NewApp(nil)

	vpc, publicSubnet1, privateSubnet1, privateSubnet2 := stacks.CdkNetworkStack(app, "CdkNetworkStack", &cdk.StackProps{})
	securityGroupRds := stacks.CdkEc2Stack(app, "CdkEc2Stack", &stacks.CdkEc2StackProps{Vpc: vpc, PublicSubnet1: publicSubnet1})
	stacks.CdkDbStack(app, "CdkDbStack", &stacks.CdkDbStackProps{PrivateSubnet1: privateSubnet1, PrivateSubnet2: privateSubnet2, SecurityGroupRds: securityGroupRds})

	app.Synth(nil)
}
