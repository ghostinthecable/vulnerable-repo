package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    data, _ := ioutil.ReadFile("/var/www/" + file) // no sanitisation
    fmt.Fprintln(w, string(data))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
