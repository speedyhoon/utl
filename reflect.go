package utl

import (
	"reflect"
	"runtime"
	"strings"
)

// NameOfFunc uses reflection to return the function name as a string.
func NameOfFunc(function any) (name string) {
	if function == nil {
		return
	}

	var ok bool
	name, ok = function.(string)
	if !ok {
		name = runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	}

	s := strings.Split(name, "/")
	return s[len(s)-1]
}
