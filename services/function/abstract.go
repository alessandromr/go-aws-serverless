package function

import (
	"time"
	"github.com/alessandromr/goserverlessclient/utils"
)

//Rollback call the object's DeleteDependencies method
func Rollback(input DeleteFunctionInput, err error) {

	/*
		Check if resource already exist:
			2019/12/01 09:25:09 ResourceConflictException: Function already exist: TestFunction status code: 409, request id: 9afc6891-29d8-4477-a0c0-06b5fd8de1b0

		If this happen rollback function must not delete the resource already existing in AWS. This already existing function is created by something or someone
		else and we dont want to modify it.
	*/

	utils.ErrLog.Println("Rollback")
	utils.ErrLog.Println(err)
	time.Sleep(utils.LongSleep * time.Millisecond)
	input.DeleteDependencies(input.GetFunctionInput())
}
