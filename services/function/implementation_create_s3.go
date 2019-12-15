package function

import (
	"github.com/alessandromr/go-aws-serverless/manager/create"
	"github.com/alessandromr/go-aws-serverless/resource/lambda/permission"
	"github.com/alessandromr/go-aws-serverless/resource/s3/notification"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/alessandromr/go-aws-serverless/utils/convert"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateDependencies create all the dependencies for S3Event
func (input S3CreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	var err error

	permission := permission.LambdaPermission{
		StatementId:  "S3Event_" + *input.S3CreateEvent.Bucket + "_" + *lambdaResult.FunctionName,
		FunctionName: *lambdaResult.FunctionArn,
		SourceArn:    "arn:aws:s3:::" + *input.S3CreateEvent.Bucket,	
		Principal:    "s3.amazonaws.com",	
		Action:       "lambda:InvokeFunction",
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&permission,
	)

	notification := notification.S3NotificationConfiguration{
		Bucket:      *input.S3CreateEvent.Bucket,
		Events:      convert.StringSlice(input.S3CreateEvent.Types),
		FunctionArn: *lambdaResult.FunctionArn,
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&notification,
	)

	//Create Resources
	err = create.ExecuteCreate()
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{})
	out["Bucket"] = *input.S3CreateEvent.Bucket
	out["StatementId"] = permission.StatementId
	return out, nil
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
