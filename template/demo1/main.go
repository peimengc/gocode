package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("500")
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Name: "裴孟",
		Age:  19,
	}
	t, _ := template.ParseFiles("./hello.tmpl")
	_ = t.Execute(w, user)
}
