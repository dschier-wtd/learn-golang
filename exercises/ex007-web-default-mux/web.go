package main

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello World!"))
}

func viewArticle(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Display an article"))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Create an article"))
}

func main() {
  http.HandleFunc("/", home)
  http.HandleFunc("/article/view", viewArticle)
  http.HandleFunc("/article/create", createArticle)

  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", nil)
  log.Fatal(err)
}
