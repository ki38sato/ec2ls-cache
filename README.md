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
|--columns|specify columns displaying (csv). use <`ec2.Instance` field> or `tag:<tagKey>` or `tagAll`. see [`ec2.Instance` field](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#Instance) (only string field). default is `"InstanceId"`|
|--filters|use filters. see [aws document](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput)|
|--profile|specify aws credential profile|
|--region|specify aws region|
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

- Use columns
```
$ ec2ls-cache -u --cachename test --columns "PrivateIpAddress,InstanceId,tag:Name,tagAll"
```


FEATURE
---
- --sortcolumn option
- cache expired limit option?
- build for windows?

LICENSE
---
MIT
