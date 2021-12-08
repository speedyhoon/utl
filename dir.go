package utl

import "os"

// Cwd returns the current working directory where the application is stored.
func Cwd() (path string) {
	var err error
	path, err = os.Getwd()
	if err != nil {
		return "."
	}

	return
}
