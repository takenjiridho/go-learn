package model

import "time"

//Profile
type Profile struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profiles []Profile

//
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
