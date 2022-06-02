package utils

// must panics when an error is thrown
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
