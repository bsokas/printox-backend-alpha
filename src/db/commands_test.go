package db

import (
  "testing"
)

func TestFindUsernameExists(t *testing.T) {
  testUsername := "test-user-1"
  if result := FindUsername(testUsername); result == "" {
    t.Fatalf("Expected temp database to have username %s\n", testUsername)
  }
}

func TestGetUsers(t *testing.T){
  expectedLen := len(TempUserMap)
  users := GetUsers()
  if len(users) != expectedLen {
    t.Fatalf("Expected %d users and got collection of %d instead\n", expectedLen, len(users))
  }
}
