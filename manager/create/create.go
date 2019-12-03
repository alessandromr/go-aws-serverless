package create

import (
	resource "github.com/alessandromr/go-serverless-client/resource"
)

//ResourcesList is a list of AWS resources ready to be created
var ResourcesList []resource.AWSResource

////ExecuteCreate will create all resources saved inside ResourcesList
//func ExecuteCreate() {
//	for _, v := range ResourcesList {
//		utils.InfoLog.Printf("Creating %T\n", v)
//		res, err := v.Create()
//		if err != nil {
//			rollback.ExecuteRollback()
//		} else {
//			rollback.ResourcesList = append(rollback.ResourcesList, res)
//		}
//	}
//}
