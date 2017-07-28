ec2ls-cache
===
`ec2ls-cache` is a simple command for describe instances with local cache.

Why cache ?
---
for MFA

CommandLine Option
---
|option key|description|
|---|---|
|--cachename|specify cachename. cache path is `~/.cache/ec2ls-cache/<cachename>`. default cachename is `out`|
|--filters|use filters. see [aws document](https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput)|
|--profile|specify aws credential profile|
|--region|specify aws region|
|--updateCache, -u|read from AWS and store cache.|

Default Output (no configurable now)
---
PrivateIP, InstanceID, TagName

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


FEATURE
---
- --columns option
- --sortcolumn option
- cache expired limit option?
- build for windows

LICENSE
---
MIT
