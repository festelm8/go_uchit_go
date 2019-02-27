package utils

// checkError check errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}