package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
  "os"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  port := os.Getenv("PORT")

  if port == "" {
    port = "3000"
  }

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":"+port)
}
