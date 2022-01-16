package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var endpoint = "http://localhost:4566"

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					SigningRegion: "ap-northeast-1",
					URL:           endpoint,
				}, nil
			}),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	output, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String(time.Now().Format("test-20060102-150405")),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", output)
}
