package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"
)

var EndpointURL = "http://localhost:4566"
var AWSRegion = "ap-northeast-1"
var BucketName = "hello"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if EndpointURL != "" {
					return aws.Endpoint{
						SigningRegion: AWSRegion,
						URL:           EndpointURL,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			}),
		),
	)
	if err != nil {
		panic(fmt.Errorf("LoadDefaultConfig failed:%v", err))
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(BucketName),
	})
	if err != nil {
		panic(fmt.Errorf("S3Client#ListObjects failed:%v", err))
	}
	for i, c := range output.Contents {
		log.Printf("[%d] %s", i, *c.Key)
	}

	key := time.Now().UTC().Format("20060102150405.000") + ".log"
	r := bytes.NewReader([]byte(key))
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    &key,
		Body:   r,
	})
	log.Printf("add %s", key)
}
