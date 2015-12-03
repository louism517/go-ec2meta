# go-ec2meta
Simple tool for querying the metadata service from an EC2 instance

## Usage

(For these examples, the binary has been compiled with `GOOS=linux GOARCH=amd64 go build -o meta`)

A check will be performed to confirm the availabilty of the EC2 metadata service (in other words - are you on an EC2 box?)

Providing the `-h` or `--help` flag lists available metadata:

```
# meta -h
usage: meta ATTRIBUTE

Valid attributes are:
ami-id
ami-launch-index
ami-manifest-path
block-device-mapping/
hostname
iam/
instance-action
instance-id
instance-type
local-hostname
local-ipv4
mac
metrics/
network/
placement/
profile
public-hostname
public-ipv4
public-keys/
reservation-id
security-groups
services/
region
```

Then quite simply:

```
# meta region
us-east-1
# meta instance-type
m3.xlarge
```
