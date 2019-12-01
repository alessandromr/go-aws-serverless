package utils

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

var (
	//InfoLog log to Stdout informations
	InfoLog = log.New(os.Stdout, "[INFO] ", log.Ltime)

	//WarnLog log to Stdout warning
	WarnLog = log.New(os.Stdout, "[WARN] ", log.Ltime)

	//ErrLog log to Stderr errors
	ErrLog = log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Lshortfile)
)

//CheckErr give an exception if the given error is not nil
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//CheckAWSErrExpect404 check the given error but let pass all 404 (ResourceNotFound) and simple log the string parameter
func CheckAWSErrExpect404(err error, resourceName string) {
	if err != nil {
		awsErr, cerr := err.(awserr.Error)
		//if is an aws error check aws logic
		if cerr == true{
			switch awsErr.Code() {
			case "ResourceNotFoundException":
				WarnLog.Printf("%s: resource not found\n", resourceName)
				break
			case  "NotFoundException":
				WarnLog.Printf("%s: resource not found\n", resourceName)
				break
			default:
				log.Fatal(err)
				break
			}
		} else {
			//else simply exit
			log.Fatal(err)
		}
	}
}
