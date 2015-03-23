package main

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<h1>hello world</h1>")
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.ListenAndServe(":5000", nil)
}
