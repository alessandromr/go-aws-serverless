package function

type S3Event struct {
	Bucket string
	Prefix string
	Suffic string
	Key    string
}

type HTTPEvent struct {
	Path   string
	Method string
}
