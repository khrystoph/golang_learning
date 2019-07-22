//Package s3fetcher is a test program to see errors returned on bucket issues.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	certDir   = "certs"
	awsRegion = "us-west-2"
)

var (
	s3bucket, filePrefix, sessionProfile string
)

func init() {
	flag.StringVar(&s3bucket, "bucket", "moreip.jbecomputersolutions.com", "Enter your s3 bucket to pull from here.")
	flag.StringVar(&filePrefix, "prefix", "certs", "Enter the object prefix where you stored the certs.")
	flag.StringVar(&sessionProfile, "profile", "default", "enter the profile you wish to use to connect. Default: default")
}

func listOjbects() (objectList *s3.ListObjectsOutput, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials("", sessionProfile),
	})
	if err != nil {
		fmt.Println("Error setting up session.", err)
		os.Exit(1)
	}
	svc := s3.New(sess)
	input := &s3.ListObjectsInput{
		Bucket:  &s3bucket,
		MaxKeys: aws.Int64(2),
		Prefix:  &filePrefix,
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(awserr.Error(aerr))
		}
		return
	}

	if len(result.Contents) == 0 {
		return nil, errors.New("no ojbects found in bucket/prefix")
	}
	return result, nil
}

//syncObjects pulls (or pushes) objects to or from s3 bucket/prefix.
func syncObjects(certs *s3.ListObjectsOutput) (err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials("", sessionProfile),
	})
	if err != nil {
		return err
	}
	if _, err = os.Stat("certs"); os.IsNotExist(err) {
		err = os.Mkdir("certs", 0755)
	}
	if err != nil {
		return err
	}
	downloader := s3manager.NewDownloader(sess)
	for object := range certs.Contents {
		input := &s3.GetObjectInput{
			Bucket: &s3bucket,
			Key:    certs.Contents[object].Key,
		}
		certfile := strings.Join([]string{*certs.Contents[object].Key}, "")
		f, err := os.Create(certfile)
		if err != nil {
			return err
		}
		cert, err := downloader.Download(f, input)
		if err != nil {
			return err
		}
		fmt.Printf("Downloaded file, %d bytes\n", cert)

	}

	return nil
}

func main() {
	flag.Parse()

	certList, err := listOjbects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = syncObjects(certList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
