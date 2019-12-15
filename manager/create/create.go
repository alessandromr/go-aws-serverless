package create

import (
	"github.com/alessandromr/go-aws-serverless/manager/rollback"
	resource "github.com/alessandromr/go-aws-serverless/resource"
	"github.com/alessandromr/go-aws-serverless/utils"
	"time"
)

//ResourcesList is a list of AWS resources ready to be created
var ResourcesList []resource.AWSResource

//ExecuteCreate will create all resources saved inside ResourcesList
func ExecuteCreate() {
	time.Sleep(utils.LongSleep * time.Millisecond)
	for _, v := range ResourcesList {
		utils.InfoLog.Printf("Creating %T\n", v)
		err := v.Create()
		if err != nil {
			rollback.ExecuteRollback()
		} else {
			rollback.ResourcesList = append(rollback.ResourcesList, v)
			time.Sleep(utils.LongSleep * time.Millisecond)
		}
	}
}
