package utl

import (
	"path/filepath"
)

func CleanPath(filePath string) (_ string, err error) {
	// TODO resolve paths prefixed with %USERPROFILE%
	return filepath.Abs(filePath)
}
