package main

import (
  "fmt"
  "net/http"
)

func main() {
  var confereenceName = "Sam Li"
  fmt.Println("My Name is", confereenceName)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  http.HandleFunc("/api/blogs", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API Blogs: %s\n", r.URL.Path)
  })

  http.ListenAndServe(":80", nil)
}