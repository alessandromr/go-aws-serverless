package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func CreateFunction(input CreateFunctionInput) {
	//Create Lambda Client
	svc := lambda.New(auth.Sess)
	//Create lambda function
	lambdaConf, err := svc.CreateFunction(input.GetFunctionInput())
	utils.CheckErr(err)
	//Create Dependencies
	input.CreateDependencies(lambdaConf)
}

