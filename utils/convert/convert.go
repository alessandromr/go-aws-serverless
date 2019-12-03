package convert

//StringSlice convert a slice of pointers of string to a slice of strings
func StringSlice(slice []*string) []string {
	var retSlice []string
	for k := range slice {
		retSlice = append(retSlice, *slice[k])
	}
	return retSlice
}
