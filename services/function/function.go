package function

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/alessandromr/goserverlessclient/utils"
)

func CreateFunction(input CreateFunctionInput) {
	//Create Lambda Client
	svc := lambda.New(session.New())
	_, err = svc.CreateFunction(input.GetFunctionInput())
	utils.CheckErr(err)
}
