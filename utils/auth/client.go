package auth

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var Client AWSClient

type AWSClient struct {
	Apigatewayconn       *apigateway.APIGateway
	Apigatewayv2conn     *apigatewayv2.ApiGatewayV2
	Lambdaconn           *lambda.Lambda
	S3conn               *s3.S3
	SQSconn              *sqs.SQS
	SNSconn              *sns.SNS
	CloudwatchEventsConn *cloudwatchevents.CloudWatchEvents
}

func MakeClient(sess *session.Session) {
	Client = AWSClient{
		Apigatewayconn:       apigateway.New(sess),
		Apigatewayv2conn:     apigatewayv2.New(sess),
		Lambdaconn:           lambda.New(sess),
		S3conn:               s3.New(sess),
		SQSconn:              sqs.New(sess),
		SNSconn:              sns.New(sess),
		CloudwatchEventsConn: cloudwatchevents.New(sess),
	}
}
