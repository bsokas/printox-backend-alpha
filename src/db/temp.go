package db

// temp.go acts as a temporary database layer in RAM
// a database layer will be implemented later

type User struct {
  Username string
  Password string
  // Email string
  LocationLat float64
  LocationLong float64
}

// Temporary collection of User Profiles
var TempUserMap = map[int]User{
  1: User{"test-user-1", "test-pw-1", 36.011341, -78.927193},
  2: User{"test-user-2", "test-pw-2", 36.011590, -78.896810},
}
