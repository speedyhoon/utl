# flagvar

Provides type `StrList` that implements the `flag.Value` interface for `[]string`.

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
