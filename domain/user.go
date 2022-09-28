package domain

import "time"

type User struct {
	Id        string
	Email     string
	Password  string
	FullName  string
	ShowCount int
	Secret    string
	UniqueId  *string
	CreatedAt time.Time
}