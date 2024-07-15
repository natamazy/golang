package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	id        int
}

func New(fn string, ln string, givenId int) (*User, error) {
	if fn == "" || ln == "" || givenId < 0 {
		return nil, errors.New("invalid input for user details")
	}

	return &User{
		firstName: fn,
		lastName:  ln,
		id:        givenId,
	}, nil
}

func (u *User) PrintDetails() {
	fmt.Println(u.firstName, u.lastName, u.id)
}

func (u *User) SetFirstName(newFirstName *string) {
	u.firstName = *newFirstName
}
