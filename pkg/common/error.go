package common

// HandleError panics if err is not nil.
//
// Parameters:
//   - err: the error to handle.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
