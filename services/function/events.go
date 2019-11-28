package function

type S3CreateEvent struct {
	Bucket *string
	Prefix *string
	Suffic *string
	Types []*string
	Key    *string
	Existing bool
}

type HTTPCreateEvent struct {
	Path   *string
	Method *string
	Existing bool
	ApiId *string
	ApiName *string
	ExecutionRole *string
}

type S3DeleteEvent struct {
	Bucket *string
	StatementId *string
	ToDelete bool
}

type HTTPDeleteEvent struct {
	ApiId *string
	ResourceId *string
	Method *string
}
