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
