package main

import "fmt"

type userStorage interface {
	getUserByID(id int) (*User, error)
}
type app struct {
	storage userStorage
}

type token struct {
	userID int
}

type Profile User

func (a app) getMyProfile(token token) (*Profile, error) {
	user, err := a.storage.getUserByID(token.userID)

	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}

	var p Profile

	p = Profile(*user)

	return &p, nil
}
