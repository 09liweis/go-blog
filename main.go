package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  http.HandleFunc("/api/blogs", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API Blogs: %s\n", r.URL.Path)
  })

  http.ListenAndServe(":80", nil)
}