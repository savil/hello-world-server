package main
import (
    "fmt"
    "net/http"
    "os"
)
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(os.Stderr, "Headers: %+v\n", req.Header)

    fmt.Fprintf(w, "hello world 10\n")
}

func main() {
    http.HandleFunc("/hello", hello)

    fmt.Println("Serving requests at port 8080\n")
    http.ListenAndServe(":8080", nil)
}
