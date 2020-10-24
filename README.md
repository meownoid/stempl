# stempl

`stempl` implements simple and fast named formatting in Go programming language.

### Installation
```shell script
go get github.com/meownoid/stempl
```

### Example
```go
package main

import (
	"fmt"
	"github.com/meownoid/stempl"
)

func main() {
	result, err := stempl.Format(
		"The {{quick}} {brown} {fox} jumps over the lazy {dog}",
		map[string]string{
			"brown": "gray",
			"fox": "cat",
			"dog": "owl",
		},
	)
	if err != nil {
		panic(err)
	}
	// The {quick} gray cat jumps over the lazy owl
	fmt.Println(result)
}
```
