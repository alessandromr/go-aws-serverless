package function

func Rollback(input DeleteFunctionInput, err error) {
	//check if already exist
	input.DeleteDependencies(input.GetFunctionInput())
}
