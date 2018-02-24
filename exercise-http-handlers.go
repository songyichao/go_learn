package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str.Greeting, str.Punct, str.Who)
}
func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", Struct{"Hello", ":", "syc"})
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
