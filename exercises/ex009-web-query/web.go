package main

import (
  "log"
  "net/http"
  "strconv" // new import
  "fmt" // new import
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)

    return
  }

  w.Write([]byte("Hello World!"))
}

func viewArticle(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))

  if err != nil || id < 1 {
    http.NotFound(w, r)

    return
  }

  fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

    return
  }

  // the above code is roughly the short form / improved form of the below code
  // if r.Method != "POST" {
  //   w.Header().Set("Allow", "POST")

  //   w.WriteHeader(405)
  //   w.Write([]byte("Method not allowed"))
  //   return
  // }

  w.Write([]byte("Create a new article..."))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/article/view", viewArticle)
  mux.HandleFunc("/article/create", createArticle)

  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
