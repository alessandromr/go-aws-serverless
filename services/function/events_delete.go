package function

//S3DeleteEvent S3 trigger of a lambda function
type S3DeleteEvent struct {
	Bucket      *string
	StatementId *string
}

//HTTPDeleteEvent  HTTP (API Gateway) trigger of a lambda function
type HTTPDeleteEvent struct {
	ApiId         *string
	ResourceId    *string
	Method        *string
	ExecutionRole *string
}

//SQSDeleteEvent SQS trigger of a lambda function
type SQSDeleteEvent struct {
	Existing bool
	QueueUrl *string
}
