# Headlines

This is a simple little application that generates headlines based on random
strings contained in data files. It acts to create MabLibs-like sentences
that mimic some of the nonsense headlines we've been having since 2020 started. 

## Example

``` golang

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tpryan/headlines"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	err := headlines.LoadCache("../../data")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	h, err := headlines.NewHeadline()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	fmt.Printf("%s\n", h.Sprintln())

}
```

_This is not an official Google Product._