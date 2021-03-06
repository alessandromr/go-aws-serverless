package rollback

import (
	resource "github.com/alessandromr/go-aws-serverless/resource"
	"github.com/alessandromr/go-aws-serverless/utils"
	"time"
)

//ResourcesList is a list of AWS resources ready to be rollbacked
var ResourcesList []resource.AWSResource

//ExecuteRollback will rollback (delete) all resources saved inside ResourcesList
func ExecuteRollback() {
	time.Sleep(utils.LongSleep * time.Millisecond)
	for _, v := range ResourcesList {
		utils.ErrLog.Printf("Rollback %T\n", v)
		v.Delete()
		time.Sleep(utils.LongSleep * time.Millisecond)
	}
}
