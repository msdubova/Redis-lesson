package main

import "time"

type storage struct {
}
type User struct {
	ID     int
	Name   string
	ReadAt time.Time
}

func (s storage) getUserByID(id int) (*User, error) {
	time.Sleep(10 * time.Second)

	return &User{
		ID:     id,
		Name:   "John",
		ReadAt: time.Now(),
	}, nil
}
