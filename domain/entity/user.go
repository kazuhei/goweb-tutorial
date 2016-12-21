package entity

import "strconv"

// User ユーザーentity
type User struct {
	ID       int
	Username string
}

func (u *User) ToString() string {
	return strconv.Itoa(u.ID) + u.Username
}
