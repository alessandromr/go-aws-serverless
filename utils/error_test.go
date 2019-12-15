package utils

import (
	"testing"
)

func TestCheckErrPass(t *testing.T) {
	var err error
	CheckErr(err)
}
