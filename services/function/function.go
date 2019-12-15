package function

import (
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// CreateFunction will create function and all the dependencies
func CreateFunction(input CreateFunctionInput) (map[string]interface{}, error) {
	var err error
	//Create response Object
	out := make(map[string]interface{})

	//Create Client
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn

	//Create lambda function
	utils.InfoLog.Println("Creating The Lambda Function")
	lambdaConf, err := svc.CreateFunction(input.GetFunctionInput())
	utils.CheckErr(err)

	//Create Dependencies
	utils.InfoLog.Println("Creating The Dependencies")
	out, err = input.CreateDependencies(lambdaConf)

	if err != nil {
		//Rollback
		utils.InfoLog.Println("Deleting The Lambda Function")
		_, lerr := svc.DeleteFunction(&lambda.DeleteFunctionInput{
			FunctionName: lambdaConf.FunctionArn,
		})
		utils.CheckAWSErrExpect404(lerr, "Lambda Function")
		return nil, err
	}

	//Create Output
	out["FunctionArn"] = *lambdaConf.FunctionArn
	return out, nil
}

// DeleteFunction will delete the function and all the dependencies
func DeleteFunction(input DeleteFunctionInput) {
	//Create Lambda Client
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn
	lambdaConf := input.GetFunctionInput()

	//Delete Dependencies
	utils.InfoLog.Println("Deleting The Dependencies")
	input.DeleteDependencies(lambdaConf)

	//Delete lambda function
	utils.InfoLog.Println("Deleting The Lambda Function")
	_, err := svc.DeleteFunction(lambdaConf)
	utils.CheckAWSErrExpect404(err, "Lambda Function")
}

// ReadFunction will return the function and all the dependencies details
func ReadFunction(input ReadFunctionInput) (map[string]interface{}, error) {
	var out map[string]interface{}

	//Create Lambda Client
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn
	lambdaConf := input.GetFunctionConfiguration()

	//Read lambda function
	utils.InfoLog.Println("Reading The Lambda Function")
	funcResponse, err := svc.GetFunctionConfiguration(lambdaConf)
	if err != nil {
		return nil, err
	}

	//Read Dependencies
	utils.InfoLog.Println("Reading The Dependencies")
	out = input.ReadDependencies(funcResponse)

	out["FunctionArn"] = *funcResponse.FunctionArn
	out["Role"] = *funcResponse.Role
	out["FunctionName"] = *funcResponse.FunctionName
	out["Handler"] = *funcResponse.Handler
	out["MemorySize"] = *funcResponse.MemorySize
	out["Runtime"] = *funcResponse.Runtime
	out["Timeout"] = *funcResponse.Timeout
	out["Version"] = *funcResponse.Version
	out["LastModified"] = *funcResponse.LastModified
	out["CodeSha256"] = *funcResponse.CodeSha256
	out["CodeSize"] = *funcResponse.CodeSize
	out["Description"] = *funcResponse.Description

	// out["TracingConfig"] = *funcResponse.TracingConfig
	// out["VpcConfig"] = *funcResponse.VpcConfig
	// out["Layers"] = funcResponse.Layers
	// out["Environment"] = *funcResponse.Environment
	// out["DeadLetterConfig"] = *funcResponse.DeadLetterConfig
	return out, nil
}

// UpdateFunction will update function and all the dependencies
func UpdateFunction(input UpdateFunctionInput) (map[string]interface{}, error) {
	var err error
	//Create response Object
	out := make(map[string]interface{})

	//Create Client
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn

	//Create lambda function
	utils.InfoLog.Println("Updating The Lambda Function")
	lambdaConf, err := svc.UpdateFunctionConfiguration(input.GetUpdateFunctionConfiguration())
	utils.CheckErr(err)

	//Updating Dependencies
	utils.InfoLog.Println("Updating The Dependencies")
	input.UpdateDependencies(lambdaConf)

	//Create Output
	out["FunctionArn"] = *lambdaConf.FunctionArn
	return out, nil
}
