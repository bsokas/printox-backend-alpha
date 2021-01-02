package db

// FindUserName searches the database for a matching username
func FindUsername(username string) string {
  for _, user := range TempUserMap {
    if user.Username == username { return username }
  }

  return ""
}

// GetUsers returns the full list of Users in the temporary database
func GetUsers() []User {
  users := make([]User, len(TempUserMap))
  for k, user := range TempUserMap {
    users[k - 1] = user
  }

  return users
}
