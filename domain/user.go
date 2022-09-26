package domain

import "time"

type User struct {
	Id               string
	Email            string
	Password         string
	FullName         string
	ShowCount        int
	Secret           string
	UniqueIdentifier string
	CreatedAt        time.Time
}
