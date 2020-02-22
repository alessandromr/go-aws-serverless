package function

//S3ReadEvent S3 trigger of a lambda function
type S3ReadEvent struct {
	Bucket      *string
	StatementId *string
}

//HTTPReadEvent  HTTP (API Gateway) trigger of a lambda function
type HTTPReadEvent struct {
	ApiId         *string
	ResourceId    *string
	Method        *string
	ExecutionRole *string
}

//SQSReadEvent SQS trigger of a lambda function
type SQSReadEvent struct {
	QueueUrl *string
}
