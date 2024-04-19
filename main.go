package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world 14\n")
}

func handle500(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/500", handle500)

	fmt.Println("Serving requests at port 8080\n")
	http.ListenAndServe(":8080", nil)
}
