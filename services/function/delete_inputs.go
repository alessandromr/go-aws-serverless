package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

//DeleteFunctionInput is an interface to delete a serverless function and the relative triggger
type DeleteFunctionInput interface {
	DeleteDependencies(*lambda.DeleteFunctionInput)
	GetFunctionInput() *lambda.DeleteFunctionInput
}

//HTTPDeleteFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPDeleteFunctionInput struct {
	FunctionInput *lambda.DeleteFunctionInput
	HTTPDeleteEvent
}

func (input HTTPDeleteFunctionInput) DeleteDependencies(lambdaResult *lambda.DeleteFunctionInput) {
	svc := apigateway.New(auth.Sess)
	var err error

	//delete existing integration
	integrationInput := &apigateway.DeleteIntegrationInput{
		HttpMethod: input.HTTPDeleteEvent.Method,
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	_, err = svc.DeleteIntegration(integrationInput)
	utils.CheckErr(err)

	//delete method
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: input.HTTPDeleteEvent.Method,
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	_, err = svc.DeleteMethod(methodInput)
	utils.CheckErr(err)

	//check if resource is empty
	getResourceInput := &apigateway.GetResourceInput{
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	resourceResponse, err := svc.GetResource(getResourceInput)

	if len(resourceResponse.ResourceMethods) < 1 {
		//delete resource
		resourceInput := &apigateway.DeleteResourceInput{
			ResourceId: input.HTTPDeleteEvent.ResourceId,
			RestApiId:  input.HTTPDeleteEvent.ApiId,
		}
		_, err = svc.DeleteResource(resourceInput)
		utils.CheckErr(err)
	}

	//check if api is empty
	getResourcesInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPDeleteEvent.ApiId,
	}
	getResourcesOutput, err := svc.GetResources(getResourcesInput)

	if len(getResourcesOutput.Items) <= 1 {
		//delete api
		apiInput := &apigateway.DeleteRestApiInput{
			RestApiId: input.HTTPDeleteEvent.ApiId,
		}
		_, err = svc.DeleteRestApi(apiInput)
		utils.CheckErr(err)
	}
}

//GetFunctionInput return the DeleteFunctionInput from the custom input
func (input HTTPDeleteFunctionInput) GetFunctionInput() *lambda.DeleteFunctionInput {
	return input.FunctionInput
}

//S3DeleteFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with s3 trigger
type S3DeleteFunctionInput struct {
	FunctionInput *lambda.DeleteFunctionInput
	S3DeleteEvent
}

func (input S3DeleteFunctionInput) DeleteDependencies(lambdaResult *lambda.DeleteFunctionInput) {
	svc := s3.New(auth.Sess)
	lambdaClient := lambda.New(auth.Sess)
	var err error

	//lambda.RemovePermission
	permissionsInput := &lambda.RemovePermissionInput{
		FunctionName: lambdaResult.FunctionName,
		StatementId: input.StatementId,
	}
	_, err = lambdaClient.RemovePermission(permissionsInput)
	utils.CheckErr(err)

	//s3.CreateBucket
	if input.S3DeleteEvent.ToDelete {
		deleteBucket := &s3.DeleteBucketInput{
			Bucket: input.S3DeleteEvent.Bucket,
		}
		_, err = svc.DeleteBucket(deleteBucket)
		utils.CheckErr(err)
	}


}

//GetFunctionInput return the DeleteFunctionInput from the custom input
func (input S3DeleteFunctionInput) GetFunctionInput() *lambda.DeleteFunctionInput {
	return input.FunctionInput
}
