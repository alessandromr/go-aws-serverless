package auth

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Client AWSClient

type AWSClient struct {
	Apigatewayconn   *apigateway.APIGateway
	Apigatewayv2conn *apigatewayv2.ApiGatewayV2
	Lambdaconn       *lambda.Lambda
	S3conn           *s3.S3
}

func MakeClient(sess *session.Session) {
	Client = AWSClient{
		Apigatewayconn:   apigateway.New(sess),
		Apigatewayv2conn: apigatewayv2.New(sess),
		Lambdaconn:       lambda.New(sess),
		S3conn:           s3.New(sess),
	}
}
