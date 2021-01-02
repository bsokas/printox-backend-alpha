package users

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "db"
)

type SignupRequest struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Street string `json:"addressStreet"`
  City string `json:"addressCity"`
  State string `json:"addressState"`
  PostalCode string `json:"addressZip"`
  UserType []string `json:"userType"` // must be "printer" or "user"
}

// ParseSignUpRequest take an http Request object and maps the submission
// If there is an error parsing or the username exists, a non-nil error is returned
func ParseSignUpRequest(req *http.Request) error {
  rawBody, readErr := ioutil.ReadAll(req.Body)
  if readErr != nil {
    return readErr
  }

  var user SignupRequest
  if err := json.Unmarshal(rawBody, &user); err != nil {
    return err
  } else if db.FindUsername(user.Username) != "" {
    return fmt.Errorf("Username %s already exists\n", user.Username)
  }


  return nil
}
