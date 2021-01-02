package users

import (
  "testing"
  "encoding/json"
  "net/http"
  "bytes"
  "strings"
)

func generateTestRequest(profile SignupRequest) (*http.Request, error) {
  inBytes, marshErr := json.Marshal(profile)
  if marshErr != nil {
    return nil, marshErr
  }

  return http.NewRequest(http.MethodGet, "/test", bytes.NewReader(inBytes))
}

func TestParseSignUpRequest(t *testing.T) {
  // 401 Chapel Dr, Durham, NC 27708 -- Duke Chapel Address
  userType := []string{"printer", "user"}
  testBody := SignupRequest{"test-user", "test-pw", "401 Chapel Dr", "Durham", "NC", "27708", userType}

  req, reqErr := generateTestRequest(testBody)
  if reqErr != nil {
    t.Fatalf("%s\n", reqErr.Error())
  }

  if parseErr := ParseSignUpRequest(req); parseErr != nil {
    t.Fatalf("%s\n", parseErr.Error())
  }
}

func TestParseSignUpRequestNameInUse(t *testing.T) {
  userType := []string{"printer", "user"}
  testBody := SignupRequest{"test-user-1", "test-pw", "401 Chapel Dr", "Durham", "NC", "27708", userType}

  req, _ := generateTestRequest(testBody)
  if parseErr := ParseSignUpRequest(req); parseErr == nil {
    t.Fatalf("Expected non-nil error when passing an existing username")
  } else if strings.TrimSpace(parseErr.Error()) != "Username test-user-1 already exists" {
    t.Fatalf("Expected error declaring username to be in existence")
  }
}
