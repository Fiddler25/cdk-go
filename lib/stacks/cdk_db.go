package stacks

import (
	"github.com/Fiddler25/cdk-go/utils"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	rds "github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkDbProps struct {
	cdk.StackProps
	PrivateSubnet1   ec2.CfnSubnet
	PrivateSubnet2   ec2.CfnSubnet
	SecurityGroupRds ec2.CfnSecurityGroup
}

func CdkDb(scope constructs.Construct, id string, props *CdkDbProps) {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	subnetGroup := rds.NewCfnDBSubnetGroup(stack, jsii.String("SubnetGroup"), &rds.CfnDBSubnetGroupProps{
		DbSubnetGroupName:        jsii.String("SubnetGroup"),
		DbSubnetGroupDescription: jsii.String("SubnetGroup"),
		SubnetIds:                jsii.Strings(*props.PrivateSubnet1.Ref(), *props.PrivateSubnet2.Ref()),
	})

	rds.NewCfnDBInstance(stack, jsii.String("DbInstance"), &rds.CfnDBInstanceProps{
		DbInstanceClass:      jsii.String("db.t2.micro"),
		AllocatedStorage:     jsii.String("100"),
		DbInstanceIdentifier: jsii.String("database"),
		DbName:               jsii.String("wordpress"),
		DbSubnetGroupName:    subnetGroup.Ref(),
		Engine:               jsii.String("mysql"),
		MasterUsername:       jsii.String("root"),
		MasterUserPassword:   jsii.String(utils.EnvNames().MasterUserPassword),
		MultiAz:              true,
		Port:                 jsii.String("3306"),
		VpcSecurityGroups:    jsii.Strings(*props.SecurityGroupRds.AttrGroupId()),
	})
}
