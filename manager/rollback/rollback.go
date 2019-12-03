package rollback

import (
	"github.com/alessandromr/goserverlessclient/resources"
)

//ResourcesList is a list of AWS resources ready to be rollbacked
var ResourcesList []resources.AWSResource

//ExecuteRollback will rollback (delete) all resources saved inside ResourcesList 
func ExecuteRollback(){
	for _, v := range ResourcesList{
		utils.ErrLog.Println("Rollback ", v.(type))
		v.Delete()
	}
}