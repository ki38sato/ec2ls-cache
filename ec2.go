package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func findEc2s(profile string, region string, filters []string) ([]Ec2Info, error) {
	instances, err := findInstances(profile, region, filters)
	if err != nil {
		return nil, err
	}
	infolist := make([]Ec2Info, 0)
	for _, i := range instances {
		tagName := findTagName(i.Tags)
		infolist = append(infolist, Ec2Info{
			Name:      tagName,
			ID:        *i.InstanceId,
			PrivateIP: *i.PrivateIpAddress,
		})
	}
	return infolist, nil
}

func findInstances(profile string, region string, filters []string) ([]*ec2.Instance, error) {
	p := Params{}
	if profile != "" {
		p.profile = profile
	}
	if region != "" {
		p.region = region
	}
	sess, err := newAwsSession(p)
	if err != nil {
		return nil, err
	}

	svc := ec2.New(sess)

	// TODO: loop more than 1000
	params := &ec2.DescribeInstancesInput{}
	awsFilters, err := buildFilters(filters)
	if err != nil {
		return nil, err
	}
	if len(awsFilters) > 0 {
		params.Filters = awsFilters
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	instances := make([]*ec2.Instance, 0)
	for _, r := range resp.Reservations {
		instances = append(instances, r.Instances...)
	}

	return instances, nil
}

func findTagName(tags []*ec2.Tag) string {
	for _, t := range tags {
		if *t.Key == "Name" {
			return *t.Value
		}
	}
	return ""
}

func buildFilters(filters []string) ([]*ec2.Filter, error) {
	// filters=[]string{"Name1=Value11","Name2=Value21,Value22"}
	awsFilters := make([]*ec2.Filter, 0)
	for _, f := range filters {
		arr1 := strings.Split(f, "=")
		if len(arr1) != 2 {
			return nil, fmt.Errorf("filter:%s is invalid", f)
		}
		arr2 := strings.Split(arr1[1], ",")
		values := make([]*string, 0)
		for _, a := range arr2 {
			values = append(values, aws.String(a))
		}
		awsFilter := &ec2.Filter{
			Name:   aws.String(arr1[0]),
			Values: values,
		}

		awsFilters = append(awsFilters, awsFilter)
	}
	return awsFilters, nil
}
