ec2ls-cache
===
`ec2ls-cache` is a simple command for describe instances with local cache.

Why cache ?
---
for MFA

CommandLine Option
---
|option key|description|
|---|---|---|
|--updateCache, -u|read from AWS and store cache.|
|--cachename|specify cachename. cache path is `~/.cache/ec2ls-cache/<cachename>`. default cachename is `out`|
|--profile|specify aws credential profile|
|--region|specify aws region|

Default Output (no configurable now)
---
PrivateIP, InstanceID, TagName

Example Usage
---
- Read from AWS and store cache.
```
$ ec2ls-cache -u
```

- Read from cache
```
$ ec2ls-cache
```


FEATURE
---
- --filters option
- --columns option
- --sortcolumn option
- cache expired limit option?
- build for windows

LICENSE
---
MIT
