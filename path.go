//go:build darwin || linux

package utl

import (
	"os/user"
	"path/filepath"
	"strings"
)

func CleanPath(filePath string) (_ string, err error) {
	const homeDir = "~"
	if strings.HasPrefix(filePath, homeDir) {
		var usr *user.User
		usr, err = user.Current()
		if err != nil {
			return
		}
		filePath = strings.Replace(filePath, homeDir, usr.HomeDir, 1)
	}

	return filepath.Abs(filePath)
}
