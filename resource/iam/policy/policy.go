package role

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/iam"
)

//IamPolicy
type IamPolicy struct {
	Description    string
	Path           string
	PolicyName     string
	PolicyDocument string
	PolicyArn      string
}

//Create the given resources
func (resource *IamPolicy) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.IamConn
	cretePolicyInput := &iam.CreatePolicyInput{
		Description:    &resource.Description,
		Path:           &resource.Path,
		PolicyName:     &resource.PolicyName,
		PolicyDocument: &resource.PolicyDocument,
	}
	_, err := svc.CreatePolicy(cretePolicyInput)
	return err
}

//Delete the given resources
func (resource IamPolicy) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.IamConn
	deletePolicyInput := &iam.DeletePolicyInput{
		PolicyArn: &resource.PolicyArn,
	}
	_, err := svc.DeletePolicy(deletePolicyInput)
	return err
}
