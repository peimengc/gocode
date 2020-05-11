package main

import (
	"fmt"
	"net/http"
)

type Test struct{}

func (t *Test) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte(r.Method+"\n"))
	_, _ = rw.Write([]byte(r.URL.String()+"\n"))
}

func main() {
	http.HandleFunc("/", hello)
	http.Handle("/test", &Test{})
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println("500")
		return
	}
}

func hello(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("你好"))
}
