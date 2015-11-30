// Package ec2meta contains some functions to assist in
// the querying of the ec2 metadata service, using the offical AWS SDK
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 1 {
		fmt.Println("Please supply an attribute to query")
		os.Exit(1)
	}

	meta := ec2metadata.New(session.New())
	if meta.Available() == false {
		fmt.Println("Not an EC2 instance or Metadata service unavailable")
		os.Exit(2)
	}

	var result string
	var err error

	if strings.ToLower(os.Args[1]) == "region" {
		result, err = meta.Region()
	} else {
		result, err = meta.GetMetadata(strings.ToLower(os.Args[1]))
	}

	if err != nil {
		fmt.Println("Problem querying metadata service.")
		os.Exit(2)
	}

	fmt.Println(result)

}
