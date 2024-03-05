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

// IsDir determines if the path is a directory.
func IsDir(path string) (bool, error) {
	fs, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fs.IsDir(), nil
}
