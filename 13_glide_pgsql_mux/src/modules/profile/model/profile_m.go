package model

import (
	"fmt"
	"time"
)

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

func Newprofile() *Profile {
	fmt.Print("go model")
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}
