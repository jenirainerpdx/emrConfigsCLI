package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

func RefreshNodeDetailsFromAws() map[string]NodeDetail {

	println("brrrrriiiing - calling aws")
	nodeDetailMap := make(map[string]NodeDetail)

	sess, sessionErr := session.NewSession(
		&aws.Config{
			Region: aws.String("us-west-2")},
	)

	if sessionErr != nil {
		fmt.Printf("Unable to create an aws session in region us-west-2:  %v", sessionErr)
		os.Exit(1)
	}

	svc := ec2.New(sess)

	describeErr := svc.DescribeInstanceTypesPages(nil, func(page *ec2.DescribeInstanceTypesOutput, lastPage bool) bool {
		for _, instanceType := range page.InstanceTypes {
			node := NodeDetail{
				VCPU:   aws.Int64Value(instanceType.VCpuInfo.DefaultVCpus),
				Memory: aws.Int64Value(instanceType.MemoryInfo.SizeInMiB),
			}
			nodeDetailMap[aws.StringValue(instanceType.InstanceType)] = node
		}
		return true
	})

	if describeErr != nil {
		println("Unable to describe instance types:  %v", describeErr)
	}

	return nodeDetailMap
}
