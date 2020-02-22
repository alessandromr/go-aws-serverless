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
		Path:                     &resource.Path,
		RoleName:                 &resource.RoleName,
		PermissionsBoundary:      &resource.PermissionsBoundary,
		AssumeRolePolicyDocument: &resource.AssumeRolePolicyDocument,
		Tags:                     resource.Tags,
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
