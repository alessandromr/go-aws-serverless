package convert

import (
	"testing"
)

func TestConvertStringSlicePass(t *testing.T) {
	test := "test"
	var inputSlice []*string = []*string{
		&test,
		&test,
		&test,
		&test,
	}
	response := StringSlice(inputSlice)
	if response[0] != "test" {
		t.Error("ConvertStringSlice")
	}
}
