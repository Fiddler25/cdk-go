package stacks

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	rds "github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkDbStackProps struct {
	cdk.StackProps
	PrivateSubnet1   ec2.CfnSubnet
	PrivateSubnet2   ec2.CfnSubnet
	SecurityGroupRds ec2.CfnSecurityGroup
}

func CdkDbStack(scope constructs.Construct, id string, props *CdkDbStackProps) {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	subnetgroup := rds.NewCfnDBSubnetGroup(stack, jsii.String("SubnetGroup"), &rds.CfnDBSubnetGroupProps{
		DbSubnetGroupName:        jsii.String("subnetgroup"),
		DbSubnetGroupDescription: jsii.String("subnetgroup"),
		SubnetIds:                jsii.Strings(*props.PrivateSubnet1.Ref(), *props.PrivateSubnet2.Ref()),
	})

	rds.NewCfnDBInstance(stack, jsii.String("DbInstance"), &rds.CfnDBInstanceProps{
		DbInstanceClass:      jsii.String("db.t2.micro"),
		AllocatedStorage:     jsii.String("100"),
		DbInstanceIdentifier: jsii.String("database"),
		DbName:               jsii.String("wordpress"),
		DbSubnetGroupName:    subnetgroup.Ref(),
		Engine:               jsii.String("mysql"),
		MasterUsername:       jsii.String("root"),
		MasterUserPassword:   jsii.String(CdkEnvNames().MasterUserPassword),
		MultiAz:              true,
		Port:                 jsii.String("3306"),
		VpcSecurityGroups:    jsii.Strings(*props.SecurityGroupRds.AttrGroupId()),
	})
}