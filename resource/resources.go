package resource

//AWSResource interface
type AWSResource interface {
	Create() error
	Delete() error
}
