package common

// HandleError will handle the error by panicking
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
