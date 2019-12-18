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
func ExecuteCreate() error {
	time.Sleep(utils.LongSleep * time.Millisecond)
	for _, v := range ResourcesList {
		utils.InfoLog.Printf("Creating %T\n", v)
		err := v.Create()
		if err != nil {
			rollback.ExecuteRollback()
			return err
		} else {
			rollback.ResourcesList = append(rollback.ResourcesList, v)
			time.Sleep(utils.LongSleep * time.Millisecond)
		}
	}
	ResourcesList = []resource.AWSResource{}
	return nil
}

//ExecutePartial will create all resources saved inside ResourcesList and remove them
func ExecutePartial() error {
	time.Sleep(utils.LongSleep * time.Millisecond)
	for _, v := range ResourcesList {
		utils.InfoLog.Printf("Creating %T\n", v)
		err := v.Create()
		if err != nil {
			rollback.ExecuteRollback()
			return err
		} else {
			rollback.ResourcesList = append(rollback.ResourcesList, v)
			time.Sleep(utils.LongSleep * time.Millisecond)
		}
		ResourcesList = ResourcesList[1:]
	}
	return nil
}
