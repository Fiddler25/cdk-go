package stacks

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func CdkNetwork(scope constructs.Construct, id string, props *cdk.StackProps) (
	ec2.CfnVPC,
	ec2.CfnSubnet,
	ec2.CfnSubnet,
	ec2.CfnSubnet,
	ec2.CfnSubnet,
) {
	var sprops cdk.StackProps
	if props != nil {
		sprops = *props
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	// Vpc
	vpc := ec2.NewCfnVPC(stack, jsii.String("Vpc"), &ec2.CfnVPCProps{
		CidrBlock: jsii.String("10.0.0.0/21"),
		Tags:      &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("MyVPC")}},
	})

	// PublicSubnet
	publicSubnet1 := ec2.NewCfnSubnet(stack, jsii.String("PublicSubnet1"), &ec2.CfnSubnetProps{
		AvailabilityZone:    jsii.String("ap-northeast-1a"),
		CidrBlock:           jsii.String("10.0.0.0/24"),
		VpcId:               vpc.Ref(),
		MapPublicIpOnLaunch: true,
		Tags:                &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("PublicSubnet1")}},
	})

	publicSubnet2 := ec2.NewCfnSubnet(stack, jsii.String("PublicSubnet2"), &ec2.CfnSubnetProps{
		AvailabilityZone:    jsii.String("ap-northeast-1c"),
		CidrBlock:           jsii.String("10.0.1.0/24"),
		VpcId:               vpc.Ref(),
		MapPublicIpOnLaunch: true,
		Tags:                &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("PublicSubnet2")}},
	})

	// PrivateSubnet
	privateSubnet1 := ec2.NewCfnSubnet(stack, jsii.String("PrivateSubnet1"), &ec2.CfnSubnetProps{
		AvailabilityZone: jsii.String("ap-northeast-1a"),
		CidrBlock:        jsii.String("10.0.2.0/24"),
		VpcId:            vpc.Ref(),
		Tags:             &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("PrivateSubnet1")}},
	})

	privateSubnet2 := ec2.NewCfnSubnet(stack, jsii.String("PrivateSubnet2"), &ec2.CfnSubnetProps{
		AvailabilityZone: jsii.String("ap-northeast-1c"),
		CidrBlock:        jsii.String("10.0.3.0/24"),
		VpcId:            vpc.Ref(),
		Tags:             &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("PrivateSubnet2")}},
	})

	// InternetGateway
	igw := ec2.NewCfnInternetGateway(stack, jsii.String("InternetGateway"), &ec2.CfnInternetGatewayProps{
		Tags: &[]*cdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("MyInternetGateway")}},
	})

	ec2.NewCfnVPCGatewayAttachment(stack, jsii.String("VPCGatewayAttachment"), &ec2.CfnVPCGatewayAttachmentProps{
		VpcId:             vpc.Ref(),
		InternetGatewayId: igw.Ref(),
	})

	// RouteTable
	routeTable := ec2.NewCfnRouteTable(stack, jsii.String("RouteTable"), &ec2.CfnRouteTableProps{
		VpcId: vpc.Ref(),
	})

	ec2.NewCfnRoute(stack, jsii.String("Route"), &ec2.CfnRouteProps{
		RouteTableId:         routeTable.Ref(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId:            igw.Ref(),
	})

	ec2.NewCfnSubnetRouteTableAssociation(stack, jsii.String("RouteTableAssociation1"), &ec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: routeTable.Ref(),
		SubnetId:     publicSubnet1.Ref(),
	})

	ec2.NewCfnSubnetRouteTableAssociation(stack, jsii.String("RouteTableAssociation2"), &ec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: routeTable.Ref(),
		SubnetId:     publicSubnet2.Ref(),
	})

	return vpc, publicSubnet1, publicSubnet2, privateSubnet1, privateSubnet2
}
