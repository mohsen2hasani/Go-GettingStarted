package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.Id != 0 {
		return User{}, errors.New("new User must not include id or it must be set to zero")
	}
	user.Id = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with Id '%v' not fount", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.Id == user.Id {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with Id '%v' not fount", user.Id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with Id '%v' not fount", id)
}
