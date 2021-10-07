package brwsr

import (
	"fmt"
	"os/exec"
)

func Open(url string) (err error) {
	err = exec.Command("xdg-open", url).Start()
	if err != nil {
		return fmt.Errorf(errMsg, url, err)
	}

	return
}
