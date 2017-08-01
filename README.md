ec2ls-cache
===
`ec2ls-cache` is a simple command for describe instances with local cache.

Why cache ?
---
To avoid MFA input.

CommandLine Option
---
|option key|description|
|---|---|
|--cachename|specify cachename. cache path is `~/.cache/ec2ls-cache/<cachename>`. default cachename is `out`|
|--columns|specify columns displaying (csv). use <`ec2.Instance` field> or `Tag:<tagKey>` or `TagAll`. see [`ec2.Instance` field](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#Instance). default is `"InstanceId"`|
|--filters|use filters. see [aws document](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput)|
|--profile|specify aws credential profile|
|--region|specify aws region|
|--sortcolumn|specify sort key column|
|--updateCache, -u|read from AWS and store cache.|

Example Usage
---
- Read from AWS and store cache.
```
$ ec2ls-cache -u --cachename stage --profile yourProfile --region ap-northeast-1
```

- Read from cache
```
$ ec2ls-cache --cachename stage
```

- Use filters
```
$ ec2ls-cache -u --cachename prod --filters "instance-state-name=running" --filters "tag:Env=production"
```
  - filters depends aws-sdk-go#ec2.DescribeInstancesInput Fiters. see [aws document](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput)
  - filters sample list
    - availability-zone
    - image-id
    - instance-id
    - instance-state-code
    - instance-state-name
    - instance-type
    - ip-address
    - private-ip-address
    - subnet-id
    - tag:key=value
    - tag-key
    - tag-value
    - vpc-id
    - and more...

- Use columns
```
$ ec2ls-cache -u --cachename test --columns "PrivateIpAddress,InstanceId,Tag:Name,TagAll"
```
  - columns depends aws-sdk-go#ec2.Instance field. see [aws document](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#Instance)
  - now support string, int, float, bool, and tags
  - tags is `Tag:<tagKey>` or `TagAll`
  - columns sample list
    - ImageId
    - InstanceId
    - InstanceType
    - PrivateIpAddress
    - PublicDnsName
    - PublicIpAddress
    - SubnetId
    - VpcId
    - and more...

- Use sortcolumn
```
$ ec2ls-cache -u --columns "PrivateIpAddress,InstanceId,Tag:Name,TagAll" --sortcolumn "Tag:Name"
```
  - columns should contains sortcolumn.


FEATURE
---
- display Slices?
- cache expired limit option?
- build for windows?

LICENSE
---
MIT
