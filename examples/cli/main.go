package main

import (
	"fmt"

	"github.com/tpryan/headlines"
)

func main() {

	if err := headlines.LoadCache("../../data"); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	h, err := headlines.New()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	fmt.Printf("%s\n", h.Sprint())

}
