package main

import (
	"github.com/Fiddler25/cdk-go/lib/stacks"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := cdk.NewApp(nil)

	vpc, publicSubnet1, privateSubnet1, privateSubnet2 := stacks.CdkNetwork(app, "CdkNetwork", &cdk.StackProps{})
	securityGroupRds := stacks.CdkEc2(app, "CdkEc2", &stacks.CdkEc2Props{Vpc: vpc, PublicSubnet1: publicSubnet1})
	stacks.CdkDb(app, "CdkDb", &stacks.CdkDbProps{PrivateSubnet1: privateSubnet1, PrivateSubnet2: privateSubnet2, SecurityGroupRds: securityGroupRds})

	app.Synth(nil)
}
