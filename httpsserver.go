package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,
    "Hi, This is an example of https service in golang!")
}

func main() {
  http.HandleFunc("/", handler)
  _, err := os.Open("cert.pem")
  if err != nil {
    panic(err)
  }
  err=http.ListenAndServeTLS(":8081", "cert.pem",
    "key.pem", nil)
  fmt.Printf("%v",err)
}