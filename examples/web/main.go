package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/tpryan/headlines"
)

var tmpl *template.Template

func healthzHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	h, err := headlines.New()
	if err != nil {
		fmt.Fprintf(w, "err: %s\n", err)
	}

	tmpl.Execute(w, h)
}

func main() {
	var err error
	if err := headlines.LoadCache("../../data"); err != nil {
		log.Fatalf("err: %s\n", err)
	}

	tmpl, err = template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	http.HandleFunc("/healtz", healthzHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
