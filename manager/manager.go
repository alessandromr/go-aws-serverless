package manager

import (
	"github.com/alessandromr/go-aws-serverless/manager/create"
	"github.com/alessandromr/go-aws-serverless/manager/rollback"
	"github.com/alessandromr/go-aws-serverless/resource"
)

//Clean remove all resources from manager's singleton
func Clean() {
	create.ResourcesList = []resource.AWSResource{}
	rollback.ResourcesList = []resource.AWSResource{}
}
