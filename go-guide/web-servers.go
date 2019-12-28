package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "Hello!", r.Method)
	fmt.Println(n, err)
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
