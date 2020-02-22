package deployment

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayDeployment
type ApiGatewayDeployment struct {
	RestApiId        string
	DeploymentId     string
	StageName        string
	StageDescription string
}

//Create the given resources
func (resource *ApiGatewayDeployment) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	deploymentInput := &apigateway.CreateDeploymentInput{
		RestApiId:        &resource.RestApiId,
		StageName:        &resource.StageName,
		StageDescription: &resource.StageDescription,
	}
	response, err := svc.CreateDeployment(deploymentInput)
	resource.DeploymentId = *response.Id
	return err
}

//Delete the given resources
func (resource *ApiGatewayDeployment) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	deplymentInput := &apigateway.DeleteDeploymentInput{
		DeploymentId: &resource.DeploymentId,
		RestApiId:    &resource.RestApiId,
	}
	_, err := svc.DeleteDeployment(deplymentInput)
	return err
}
