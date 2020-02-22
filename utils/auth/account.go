package auth

import (
	"log"

	"github.com/aws/aws-sdk-go/service/sts"
)

func GetAccountID() string {
	MakeClient(Sess)
	svc := Client.StsConn

	creds, err := Sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}
	input := &sts.GetAccessKeyInfoInput{
		AccessKeyId: &creds.AccessKeyID,
	}
	response, err := svc.GetAccessKeyInfo(input)
	if err != nil {
		log.Fatal(err)
	}
	return *response.Account
}
