package namer

import "fmt"

type User struct {
	LastName  string
	FirstName string
	Age       int
}

func (u *User) UserName(uname string) string {
	return fmt.Sprintf("%s(%s)さん、%d, %s", u.LastName, u.FirstName, u.Age, uname)
}
