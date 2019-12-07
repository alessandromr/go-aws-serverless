package rollback

import (
	resource "github.com/alessandromr/go-aws-serverless/resource"
	"github.com/alessandromr/go-aws-serverless/utils"
)

//ResourcesList is a list of AWS resources ready to be rollbacked
var ResourcesList []resource.AWSResource

//ExecuteRollback will rollback (delete) all resources saved inside ResourcesList
func ExecuteRollback() {
	for _, v := range ResourcesList {
		utils.ErrLog.Printf("Rollback %T\n", v)
		v.Delete()
	}
}
