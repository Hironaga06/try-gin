package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type Users []*User
