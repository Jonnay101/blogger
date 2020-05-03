package auth

// User -
type User struct {
	userIsAdmin bool
}

// IsAdmin -
func (u *User) IsAdmin() bool {
	return u.userIsAdmin
}

// SetUserIsAdmin -
func (u *User) SetUserIsAdmin(b bool) {
	u.userIsAdmin = b
}
