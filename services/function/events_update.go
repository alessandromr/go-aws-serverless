package function

//S3UpdateEvent S3 trigger of a lambda function
type S3UpdateEvent struct {
	Bucket      *string
	Prefix      *string
	Suffix      *string
	Types       []*string
	Key         *string
	StatementId *string
}

//HTTPUpdateEvent HTTP (API Gateway) trigger of a lambda function
type HTTPUpdateEvent struct {
	Path              *string
	Method            *string
	Existing          bool
	ApiId             *string
	ResourceId        *string
	ApiName           *string
	ExecutionRoleName *string
}
