package server

import (
  "net/http"
  "io"
  "users"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Welcome to the Printox alpha API\n")
}

// SignUpHandler receives a signup request
func SignUpHandler(w http.ResponseWriter, req *http.Request) {
  if parseErr := users.ParseSignUpRequest(req); parseErr != nil {
    w.WriteHeader(http.StatusConflict)
    io.WriteString(w, parseErr.Error())
    return
  }

  w.WriteHeader(http.StatusOK)
}

// LoginHandler receives login credentials and confirms or denies access
func LoginHandler(w http.ResponseWriter, req *http.Request) {}
