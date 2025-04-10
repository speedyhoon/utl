# flagvar 🏁

[![Go Reference](https://pkg.go.dev/badge/speedyhoon/utl/flagvar.svg)](https://pkg.go.dev/speedyhoon/utl/flagvar)
[![go report card](https://goreportcard.com/badge/github.com/speedyhoon/utl/flagvar)](https://goreportcard.com/report/github.com/speedyhoon/utl/flagvar)

Provides types that implement the `flag.Value` interface for built-in types to use with `flag.Var`, including:
* `StrList []string`
* `Uint8`
* `Uint16`

Example usage:
```go
package main

import (
	"flag"
	"fmt"
	"github.com/speedyhoon/utl/flagvar"
)

func main() {
	var types flagvar.StrList
	flag.Var(&types, "animals", "Comma separated list of animals")
	flag.Parse()

	fmt.Println(types)
	fmt.Println(flag.Args())
}
```

Run with:
```shell
go run main.go -animals cow,pig farm domestic
```

Output:
```
animals: ["cow", "pig"]
args: ["farm", "domestic"]
```
