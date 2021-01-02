package server

import (
  "net/http"
  "fmt"
  "log"
)

const DefaultPortNumber = 8013

func Start(){
  prepHandlers()

  portStr := fmt.Sprintf(":%d", DefaultPortNumber)
  log.Fatal(http.ListenAndServe(portStr, nil))
}

// private functions

func prepHandlers(){
  http.HandleFunc("/", RootHandler)
}
