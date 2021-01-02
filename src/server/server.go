package server

import (
  "net/http"
  "fmt"
  "log"
)

/*
curl -d '{"username": "test-user-1", "password": "password-1", "addressStreet": "309 Clark Street", "addressCity": "Durham", "addressState": "NC", "addressZip": "27705", "userType": ["printer", "user"]}' -H "Content-Type: application/json" -X POST localhost:8013/signup -i
*/

const DefaultPortNumber = 8013

// Start sets handlers and listens on the default port
func Start(){
  prepHandlers()

  portStr := fmt.Sprintf(":%d", DefaultPortNumber)
  log.Fatal(http.ListenAndServe(portStr, nil))
}

// private functions

// prepHandlers sets handler functions for paths
func prepHandlers(){
  http.HandleFunc("/", RootHandler)
  http.HandleFunc("/signup", SignUpHandler)
}
