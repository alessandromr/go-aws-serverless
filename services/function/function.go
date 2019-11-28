package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func CreateFunction(input CreateFunctionInput) map[string]interface{} {
	//Create response Object
	out := make(map[string]interface{})

	//Create Lambda Client
	svc := lambda.New(auth.Sess)
	//Create lambda function
	lambdaConf, err := svc.CreateFunction(input.GetFunctionInput())
	utils.CheckErr(err)

	//Create Dependencies
	out = input.CreateDependencies(lambdaConf)
	out["FunctionArn"] = *lambdaConf.FunctionArn
	return out
}

func DeleteFunction(input DeleteFunctionInput) {
	//Create Lambda Client
	svc := lambda.New(auth.Sess)
	lambdaConf := input.GetFunctionInput()

	//Delete Dependencies
	input.DeleteDependencies(lambdaConf)

	//Delete lambda function
	_, err := svc.DeleteFunction(lambdaConf)
	utils.CheckErr(err)
}
