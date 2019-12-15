package rule

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

//CloudwatchRule
type CloudwatchRule struct {
	Name               *string
	Description        *string
	RoleArn            *string
	ScheduleExpression *string
	State              *string
}

//Delete the given resources
func (resource CloudwatchRule) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.CloudwatchEventsConn

	deleteRuleInput := &cloudwatchevents.DeleteRuleInput{
		Name: resource.Name,
	}
	_, err := svc.DeleteRule(deleteRuleInput)
	return err
}
