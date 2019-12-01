package function

//S3CreateEvent S3 trigger of a lambda function
type S3CreateEvent struct {
	Bucket *string
	Prefix *string
	Suffic *string
	Types []*string
	Key    *string
}

//HTTPCreateEvent HTTP (API Gateway) trigger of a lambda function
type HTTPCreateEvent struct {
	Path   *string
	Method *string
	Existing bool
	ApiId *string
	ApiName *string
	ExecutionRole *string
}

//S3DeleteEvent S3 trigger of a lambda function
type S3DeleteEvent struct {
	Bucket *string
	StatementId *string
}

//HTTPDeleteEvent  HTTP (API Gateway) trigger of a lambda function
type HTTPDeleteEvent struct {
	ApiId *string
	ResourceId *string
	Method *string
}

//S3ReadEvent S3 trigger of a lambda function
type S3ReadEvent struct {
	Bucket *string
	StatementId *string
}

//HTTPReadEvent  HTTP (API Gateway) trigger of a lambda function
type HTTPReadEvent struct {
	ApiId *string
	ResourceId *string
	Method *string
}

