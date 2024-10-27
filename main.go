package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func main() {
	// Parse command-line flags
	zoneIDPtr := flag.String("zoneID", "", "Hosted Zone ID")
	flag.Parse()

	// Validate zoneID flag
	if *zoneIDPtr == "" {
		fmt.Println("Please provide a valid hosted zone ID using the -zoneID flag")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		panic("failed to load AWS config")
	}

	client := route53.NewFromConfig(cfg)

	resp, err := client.ListResourceRecordSets(context.TODO(), &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(*zoneIDPtr),
	})
	if err != nil {
		panic(err)
	}

	// Marshal the record sets into indented JSON
	recordSetsJSON, err := json.MarshalIndent(resp.ResourceRecordSets, "", "  ")
	if err != nil {
		panic(err)
	}

	// Print the record sets JSON to stdout
	os.Stdout.Write(recordSetsJSON)
}
