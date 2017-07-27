package main

import "github.com/aws/aws-sdk-go/service/ec2"

func findEc2s(profile string, region string) ([]Ec2Info, error) {
	instances, err := findInstances(profile, region)
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

func findInstances(profile string, region string) ([]*ec2.Instance, error) {
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

	// TODO: filters
	// TODO: loop more than 1000
	params := &ec2.DescribeInstancesInput{}
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
