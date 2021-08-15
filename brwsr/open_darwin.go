package brwsr

import "fmt"

func Open(url string) (err error) {
	err = exec.Command("open", url).Start()
	if err != nil {
		return fmt.Errorf(errMsg, url, err)
	}

	return
}
