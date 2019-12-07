package auth

import (
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Sess *session.Session
var Region string

func StartSessionWithShared(region string, profile string) {
	tmp, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewSharedCredentials("", profile),
		Region:      aws.String(region),
	})
	utils.CheckErr(err)
	Sess = tmp
}

func StartSession(region string) {
	tmp, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	utils.CheckErr(err)
	Sess = tmp
}

func SetRegion(region string) {
	Region = region
}
