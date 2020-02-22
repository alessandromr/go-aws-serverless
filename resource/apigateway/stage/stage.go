package stage

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayStage
type ApiGatewayStage struct {
	RestApiId    string
	StageName    string
	DeploymentID string
}

//Create the given resources
func (resource *ApiGatewayStage) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	stageInput := &apigateway.CreateStageInput{
		RestApiId:    &resource.RestApiId,
		DeploymentId: &resource.DeploymentID,
		StageName:    &resource.StageName,
	}
	_, err := svc.CreateStage(stageInput)
	return err
}

//Delete the given resources
func (resource *ApiGatewayStage) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	stageInput := &apigateway.DeleteStageInput{
		RestApiId: &resource.RestApiId,
		StageName: &resource.StageName,
	}
	_, err := svc.DeleteStage(stageInput)
	return err
}
