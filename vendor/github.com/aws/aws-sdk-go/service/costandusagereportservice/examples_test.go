// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package costandusagereportservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCostandUsageReportService_DeleteReportDefinition() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := costandusagereportservice.New(sess)

	params := &costandusagereportservice.DeleteReportDefinitionInput{
		ReportName: aws.String("ReportName"),
	}
	resp, err := svc.DeleteReportDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCostandUsageReportService_DescribeReportDefinitions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := costandusagereportservice.New(sess)

	params := &costandusagereportservice.DescribeReportDefinitionsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("GenericString"),
	}
	resp, err := svc.DescribeReportDefinitions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCostandUsageReportService_PutReportDefinition() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := costandusagereportservice.New(sess)

	params := &costandusagereportservice.PutReportDefinitionInput{
		ReportDefinition: &costandusagereportservice.ReportDefinition{ // Required
			AdditionalSchemaElements: []*string{ // Required
				aws.String("SchemaElement"), // Required
				// More values...
			},
			Compression: aws.String("CompressionFormat"), // Required
			Format:      aws.String("ReportFormat"),      // Required
			ReportName:  aws.String("ReportName"),        // Required
			S3Bucket:    aws.String("S3Bucket"),          // Required
			S3Prefix:    aws.String("S3Prefix"),          // Required
			S3Region:    aws.String("AWSRegion"),         // Required
			TimeUnit:    aws.String("TimeUnit"),          // Required
			AdditionalArtifacts: []*string{
				aws.String("AdditionalArtifact"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.PutReportDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}