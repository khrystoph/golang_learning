package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	//Trace is log handling for Trace level messages
	Trace *log.Logger
	//Info is log handling for Info level messaging
	Info *log.Logger
	//Warning is log handling for Warning level messaging
	Warning *log.Logger
	//Error is log handling for Error level messaging
	Error         *log.Logger
	traceHandle   io.Writer
	infoHandle    io.Writer = os.Stdout
	warningHandle io.Writer = os.Stderr
	errorHandle   io.Writer = os.Stderr

	s3Bucket, s3BucketKey string
	presignTime           = 10080 //equals 1 week of time
)

func init() {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	flag.StringVar(&s3Bucket, "bucket", "bucket.example.com", "Enter your s3 bucket to pull from here.")
	flag.StringVar(&s3Bucket, "b", "bucket.example.com", "Enter your s3 bucket to pull from here.")
	flag.StringVar(&s3BucketKey, "key", "somekey", "Enter the s3 Key of the object in the bucket.")
	flag.StringVar(&s3BucketKey, "k", "somekey", "Enter the s3 Key of the object in the bucket.")
	flag.IntVar(&presignTime, "t", presignTime, "Enter an integer in minutes that represents the time the URL is valid.")
}

//calls SDK to generate a presigned URL based on the input key
func getPresignedURL() (presignedURL string, err error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3BucketKey),
	})
	presignedURL, err = req.Presign(time.Duration(presignTime) * time.Minute)

	if err != nil {
		Error.Println("Failed to sign request", err)
		return "", err
	}

	return presignedURL, nil
}

//This program spits out pre-signed URLs so that you can send them to a user and download objects from the bucket
func main() {
	flag.Parse()

	urlStr, err := getPresignedURL()
	if err != nil {
		Error.Println("Failed to sign request", err)
		return
	}

	Info.Println(urlStr)
}
