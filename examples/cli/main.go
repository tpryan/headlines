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

	r, err := headlines.NewHeadline()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	fmt.Printf("Today, %s in %s %s %s.\n", r.Subject, r.Location, r.Verb, r.Object)

}
