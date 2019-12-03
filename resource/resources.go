package resource

//AWSResource interface
type AWSResource interface {
	//Create() (AWSResource, error)
	Delete() error
}
