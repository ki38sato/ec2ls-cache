package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func findEc2s(profile string, region string, filters []string, columns string, sortcolumn string) (map[string]interface{}, error) {
	instances, err := findInstances(profile, region, filters)
	if err != nil {
		return nil, err
	}

	// convert ec2.Instance to map[string]interface{}
	infolist := make([]map[string]interface{}, 0)

	for _, i := range instances {
		info := make(map[string]interface{})
		for _, c := range strings.Split(columns, ",") {
			cv, err := findColumnValue(i, c)
			if err != nil {
				return nil, err
			}
			info[c] = cv
		}
		infolist = append(infolist, info)
	}

	if sortcolumn != "" {
		sort.Slice(infolist, func(i, j int) bool {
			return infolist[i][sortcolumn].(string) < infolist[j][sortcolumn].(string)
		})
	}

	// set cacheinfo
	result := make(map[string]interface{})
	result["instances"] = infolist
	result["columns"] = columns

	return result, nil
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

func findColumnValue(i *ec2.Instance, columnName string) (string, error) {
	var columnValue string
	if columnName == "TagAll" {
		columnValue = findTagAll(i.Tags)
	} else if strings.Index(columnName, "Tag:") == 0 {
		columnValue = findTagValue(columnName, i.Tags)
	} else {
		r := reflect.ValueOf(i)
		f := reflect.Indirect(r).FieldByName(columnName)
		// check column value type
		if err := validateColumnType(f, columnName); err != nil {
			return "", err
		}
		columnValue = fmt.Sprintf("%v", f.Elem())
	}

	return columnValue, nil
}

func findTagAll(tags []*ec2.Tag) string {
	taglist := make([]string, 0)
	for _, t := range tags {
		taginfo := *t.Key + ":" + *t.Value
		if taginfo != "" {
			taglist = append(taglist, taginfo)
		}
	}
	return strings.Join(taglist, ";")
}

func findTagValue(columnName string, tags []*ec2.Tag) string {
	tagKey := strings.Split(columnName, ":")[1]
	for _, t := range tags {
		if *t.Key == tagKey {
			return *t.Value
		}
	}
	return ""
}

func validateColumnType(f reflect.Value, columnName string) error {
	invalidTypeList := []reflect.Kind{
		reflect.Invalid,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Array,
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Slice,
		reflect.Struct,
		reflect.UnsafePointer,
	}

	for _, t := range invalidTypeList {
		if f.Kind() == t {
			return fmt.Errorf("column: %s is invalid type: %v", columnName, t)
		}
	}
	for _, t := range invalidTypeList {
		if f.Elem().Kind() == t {
			return fmt.Errorf("column: %s is invalid type: %v", columnName, t)
		}
	}
	if f.Elem().Kind() == reflect.Ptr {
		return fmt.Errorf("column: %s is invalid type: %v", columnName, reflect.Ptr)
	}

	return nil
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
