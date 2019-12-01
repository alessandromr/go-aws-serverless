package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// CreateFunction will create function and all the dependencies
func CreateFunction(input CreateFunctionInput) map[string]interface{} {
	var err error
	//Create response Object
	out := make(map[string]interface{})

	//Create Lambda Client
	svc := lambda.New(auth.Sess)

	//Create lambda function
	utils.InfoLog.Println("Creating The Lambda Function")
	lambdaConf, err := svc.CreateFunction(input.GetFunctionInput())
	utils.CheckErr(err)

	//Create Dependencies
	utils.InfoLog.Println("Creating The Dependencies")
	out, err = input.CreateDependencies(lambdaConf)

	if err != nil{
		//Rollback
		utils.InfoLog.Println("Deleting The Lambda Function")
		_, lerr := svc.DeleteFunction(&lambda.DeleteFunctionInput{
			FunctionName: lambdaConf.FunctionArn,
		})
		utils.CheckAWSErrExpect404(lerr, "Lambda Function")
		return nil
	}

	//Create Output
	out["FunctionArn"] = *lambdaConf.FunctionArn
	return out
}

// DeleteFunction will delete the function and all the dependencies
func DeleteFunction(input DeleteFunctionInput) {
	//Create Lambda Client
	svc := lambda.New(auth.Sess)
	lambdaConf := input.GetFunctionInput()

	//Delete Dependencies
	utils.InfoLog.Println("Deleting The Dependencies")
	input.DeleteDependencies(lambdaConf)

	//Delete lambda function
	utils.InfoLog.Println("Deleting The Lambda Function")
	_, err := svc.DeleteFunction(lambdaConf)
	utils.CheckAWSErrExpect404(err, "Lambda Function")
}
