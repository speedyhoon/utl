package utl

import "strings"

const (
	GoExt        = ".go"
	GoTestSuffix = "_test" + GoExt
)

func IsGoFileName(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), GoExt)
}

func IsGoTestFileName(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), GoTestSuffix)
}

// GoTestFileName changes a path's file extension from `.go` to `_test.go`.
func GoTestFileName(goFileNamePath string) string {
	return strings.TrimSuffix(goFileNamePath, GoExt) + GoTestSuffix
}
