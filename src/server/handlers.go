package server

import (
  "net/http"
  "io"
)

func RootHandler(writer http.ResponseWriter, req *http.Request) {
  io.WriteString(writer, "Welcome to the Printox alpha API\n")
}
