package utl

import "strings"

const goExt = ".go"
const testSuffix = "_test" + goExt

func IsGoFileName(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), goExt)
}

func IsGoTestFileName(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), testSuffix)
}
