package blog

// User -
type User struct {
	admin bool
}

// IsAdmin -
func (u *User) IsAdmin() bool {
	return u.admin
}
