package utils

import "log"

//CheckErr give an exception if the given error is not nil
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
