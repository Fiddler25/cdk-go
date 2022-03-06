package stacks

import (
	"encoding/base64"
	"os"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkEc2StackProps struct {
	cdk.StackProps
	Vpc           ec2.CfnVPC
	PublicSubnet1 ec2.CfnSubnet
}

func CdkEc2Stack(scope constructs.Construct, id string, props *CdkEc2StackProps) ec2.CfnSecurityGroup {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	// SecurityGroup for EC2
	webSg := ec2.NewCfnSecurityGroup(stack, jsii.String("SecurityGroupEc2"), &ec2.CfnSecurityGroupProps{
		GroupName:        jsii.String("Web-SG"),
		GroupDescription: jsii.String("for Web"),
		VpcId:            props.Vpc.Ref(),
		SecurityGroupIngress: &[]*ec2.CfnSecurityGroup_IngressProperty{
			{
				IpProtocol: jsii.String("tcp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(22),
				ToPort:     jsii.Number(22),
			},
			{
				IpProtocol: jsii.String("tcp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(80),
				ToPort:     jsii.Number(80),
			},
		},
	})

	// SecurityGroup for RDS
	rdsSg := ec2.NewCfnSecurityGroup(stack, jsii.String("SecurityGroupRds"), &ec2.CfnSecurityGroupProps{
		GroupName:        jsii.String("Rds-SG-1"),
		GroupDescription: jsii.String("for Rds"),
		VpcId:            props.Vpc.Ref(),
		SecurityGroupIngress: &[]*ec2.CfnSecurityGroup_IngressProperty{
			{
				IpProtocol:            jsii.String("tcp"),
				FromPort:              jsii.Number(3306),
				ToPort:                jsii.Number(3306),
				SourceSecurityGroupId: webSg.AttrGroupId(),
			},
		},
	})

	// Instance
	ec2.NewCfnInstance(stack, jsii.String("Ec2Instance1"), &ec2.CfnInstanceProps{
		ImageId:          jsii.String("ami-03d79d440297083e3"),
		InstanceType:     jsii.String("t2.micro"),
		SubnetId:         props.PublicSubnet1.Ref(),
		SecurityGroupIds: jsii.Strings(*webSg.AttrGroupId()),
		KeyName:          jsii.String(CdkEnvNames().KeyName),
		UserData:         jsii.String(getUserData()),
		Tags:             &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("WebServer1")}},
	})

	return rdsSg
}

func getUserData() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.ReadFile(dir + "/workspace/App/cdk-go/bin/script/user_data.sh")
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(f)
}
