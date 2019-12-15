package function

//S3CreateEvent S3 trigger of a lambda function
type S3CreateEvent struct {
	Bucket *string
	Prefix *string
	Suffix *string
	Types  []*string
	Key    *string
}

//HTTPCreateEvent HTTP (API Gateway) trigger of a lambda function
type HTTPCreateEvent struct {
	Path          *string
	Method        *string
	Existing      bool
	ApiId         *string
	ApiName       *string
	ExecutionRole *string
}
