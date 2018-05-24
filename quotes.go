package utl

import (
	"fmt"
	"strings"
)

//AddQuotes returns value with or without surrounding single or double quote characters suitable for a [[//dev.w3.org/html5/html-author/#attributes][HTML5 attribute]] value.
func AddQuotes(in interface{}) string {
	value := strings.Replace(fmt.Sprintf("%v", in), `&`, "&amp;", -1) //will destroy any existing escaped characters like &#62;
	double := strings.Count(value, `"`)
	single := strings.Count(value, `'`)
	if single > 0 && single >= double {
		return `"` + strings.Replace(value, `"`, "&#34;", -1) + `"`
	}
	//Space, double quote, accent, equals, less-than sign, greater-than sign.
	if double > 0 || strings.ContainsAny(value, " \"`=<>") {
		return `'` + strings.Replace(value, `'`, "&#39;", -1) + `'`
	}
	return value
}
