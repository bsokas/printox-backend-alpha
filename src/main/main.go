package main

import (
  "fmt"
  // Local packages
  "server"
)

func main() {
  fmt.Println("Starting Printox alpha API")
  server.Start()
}
