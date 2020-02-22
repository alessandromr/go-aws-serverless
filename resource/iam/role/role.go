package role

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/iam"
)

//IamRole
type IamRole struct {
	Description              string
	Path                     string
	PermissionsBoundary      string
	RoleName                 string
	Tags                     []*iam.Tag
	AssumeRolePolicyDocument string
}

//Create the given resources
func (resource *IamRole) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.IamConn
	createRoleInput := &iam.CreateRoleInput{
		Description:              &resource.Description,
		RoleName:                 &resource.RoleName,
		AssumeRolePolicyDocument: &resource.AssumeRolePolicyDocument,
		Tags:                     resource.Tags,
	}
	if len(resource.Path) >= 1 {
		createRoleInput.Path = &resource.Path
	}
	if len(resource.PermissionsBoundary) >= 20 {
		createRoleInput.PermissionsBoundary = &resource.PermissionsBoundary
	}
	_, err := svc.CreateRole(createRoleInput)
	return err
}

//Delete the given resources
func (resource IamRole) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.IamConn
	deleteRoleInput := &iam.DeleteRoleInput{
		RoleName: &resource.RoleName,
	}
	_, err := svc.DeleteRole(deleteRoleInput)
	return err
}
