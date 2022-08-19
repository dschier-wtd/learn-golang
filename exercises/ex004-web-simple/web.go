package main

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World!"))
}

func main() {
  // Use the http.NewServeMux() function to initialize a new servemux, then
  // register the home function as the handler for the "/" URL pattern.
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)

  // Use the http.ListenAndServe() function to start the server.
  // If http.ListenAndServe() returns an error we use the log.Fatal() function
  // to log the error message and exit.
  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
