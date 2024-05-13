package namer

import "fmt"

type User struct {
	LastName  string
	FirstName string
	Age       int
}

// func (u User) String() string {
// 	return fmt.Sprintf("%s %s %d", u.LastName, u.FirstName, u.Age)
// }

func (u *User) UserName() string {
	return fmt.Sprintf("%s%s,%d歳", u.LastName, u.FirstName, u.Age)
}
