package utl

import (
	"os"
	"path/filepath"
)

//Cwd returns the current working directory where the application is stored.
func Cwd() string {
	path, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(path)
}
