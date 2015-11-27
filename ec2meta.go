// Package ec2meta contains some functions to assist in
// the querying of the ec2 metadata service, using the offical AWS SDK
package ec2meta

import (
  "github.com/aws/aws-sdk-go/aws/ec2metadata"
  "github.com/aws/aws-sdk-go/aws/session"
  "fmt"
)

func QueryMetadata() {
  meta := ec2metadata.New(session.New())
  localip, err := meta.GetMetadata("local-ipv4")
  if err != nil {
    fmt.Println("Cannee get the IP")
    return
  }
  fmt.Println(localip)
}
