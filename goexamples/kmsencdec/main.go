package main

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

var EndpointURL = "http://localhost:4566"
var AWSRegion = "ap-northeast-1"
var KMSKey = "alias/local-kms-key"

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
		log.Fatalf("LoadDefaultConfig failed:%v", err)
	}
	client := kms.NewFromConfig(cfg)

	// encrypt
	plainText := "test"
	encOutput, err := client.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:     aws.String(KMSKey),
		Plaintext: []byte(plainText),
	})
	if err != nil {
		log.Fatalf("kms.Encrypt failed:%v", err)
	}
	base64encrypted := base64.StdEncoding.EncodeToString(encOutput.CiphertextBlob)
	log.Printf("encrypted(base64): %s", base64encrypted)

	// decrypt
	encrypted, _ := base64.StdEncoding.DecodeString(base64encrypted)
	decOutput, err := client.Decrypt(context.TODO(), &kms.DecryptInput{
		KeyId:          aws.String(KMSKey),
		CiphertextBlob: encrypted,
	})
	if err != nil {
		log.Fatalf("kms.Decrypt failed:%v", err)
	}
	log.Printf("decrypted: %s", string(decOutput.Plaintext))
}
