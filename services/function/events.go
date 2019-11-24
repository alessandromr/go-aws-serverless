package function

type S3Event struct {
	Bucket string
	Prefix string
	Suffic string
	Types []string
	Key    string
	Existing bool
}

type HTTPEvent struct {
	Path   string
	Method string
	Existing bool
	ApiId string
	ApiName string
}
