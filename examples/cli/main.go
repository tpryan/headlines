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
