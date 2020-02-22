package auth

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sts"
)

var Client AWSClient

type AWSClient struct {
	ApigatewayConn       *apigateway.APIGateway
	Apigatewayv2Conn     *apigatewayv2.ApiGatewayV2
	LambdaConn           *lambda.Lambda
	S3Conn               *s3.S3
	SQSConn              *sqs.SQS
	SNSConn              *sns.SNS
	CloudwatchEventsConn *cloudwatchevents.CloudWatchEvents
	IamConn              *iam.IAM
	StsConn              *sts.STS
}

func MakeClient(sess *session.Session) {
	Client = AWSClient{
		ApigatewayConn:       apigateway.New(sess),
		Apigatewayv2Conn:     apigatewayv2.New(sess),
		LambdaConn:           lambda.New(sess),
		S3Conn:               s3.New(sess),
		SQSConn:              sqs.New(sess),
		SNSConn:              sns.New(sess),
		CloudwatchEventsConn: cloudwatchevents.New(sess),
		IamConn:              iam.New(sess),
		StsConn:              sts.New(sess),
	}
}
