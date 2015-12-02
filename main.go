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

func showHelp(meta *ec2metadata.EC2Metadata){

	fmt.Printf("usage: meta ATTRIBUTE\n\nValid attributes are:\n")
	categories, _ := meta.GetMetadata("/")
	fmt.Println(categories)
	fmt.Println("region")
}

func main() {

	meta := ec2metadata.New(session.New())
	if meta.Available() == false {
		fmt.Println("Not an EC2 instance or Metadata service unavailable")
		os.Exit(2)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please supply an attribute to query or -h for help")
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		showHelp(meta)
		os.Exit(0)
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
