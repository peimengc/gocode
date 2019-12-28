package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, s.Greeting+s.Punct+s.Who)
}

func main() {
	http.Handle("/string", String("你好"))
	http.Handle("/struct", Struct{"a", ":", "c"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
