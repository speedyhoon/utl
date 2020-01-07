package utl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

//Exec executes an external command & returns stdout & stderr
//src = content passed to stdin
//dir = which directory to execute the command in
//command = the command to execute
func Exec(src []byte, dir, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	var stderr io.ReadCloser
	stderr, err = cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if len(src) > 0 {
		var stdin io.WriteCloser
		stdin, err = cmd.StdinPipe()
		if err != nil {
			return nil, err
		}

		//pass src to stdin
		go func() {
			defer func() {
				if err = stdin.Close(); err != nil {
					log.Println(err)
				}
			}()

			//Ignore len(src) written to src
			if _, err = stdin.Write(src); err != nil {
				log.Println(err)
			}
		}()
	}

	var errs []string
	var output []byte
	scanOut := bufio.NewScanner(stdout)
	go func() {
		//Gather output from external command
		for scanOut.Scan() {
			output = append(output, []byte(fmt.Sprintf("%v\n", scanOut.Text()))...)
		}
		//Collect stdout scanner error
		if err = scanOut.Err(); err != nil {
			errs = append(errs, "stdout scan err: "+err.Error())
		}
	}()

	scanErr := bufio.NewScanner(stderr)
	go func() {
		//Collect all errors returned from stderr and scanner errors
		for scanErr.Scan() {
			errs = append(errs, scanErr.Text())
		}
		if err = scanErr.Err(); err != nil {
			errs = append(errs, "stderr scan err: "+err.Error())
		}
	}()

	if err = cmd.Start(); err != nil {
		errs = append(errs, err.Error())
	}

	if err = cmd.Wait(); err != nil {
		errs = append(errs, err.Error())
	}

	//Return all the errors from stdout, stderr, start & wait
	if len(errs) > 0 {
		return output, fmt.Errorf(strings.Join(errs, "\n"))
	}
	return output, nil
}

func ExecArgs(stdIn []byte, dir string, args ...string) ([]byte, error) {
	switch len(args) {
	case 0:
		return nil, errors.New("no arguments supplied")
	case 1:
		return Exec(stdIn, dir, args[0])
	}
	return Exec(stdIn, dir, args[0], args[1:]...)
}
