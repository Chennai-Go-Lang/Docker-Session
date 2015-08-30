package main

import (
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
  "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
  resp, _ := http.Post("http://192.168.0.35:8080/postcrc", "text/plain", bytes.NewBufferString("your string"))

  htmlData, err := ioutil.ReadAll(resp.Body)

  if err != nil { log.Print(err) }
  fmt.Fprintf(w, "Posted data and received: %s", htmlData)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
