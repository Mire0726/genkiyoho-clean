package model

import (
	"time"
)

type User struct {
    ID        string
    AuthToken string
    Email     string
    Password  string
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewUser(id, authToken, email, password, name string, createdAt, updatedAt time.Time) (*User, error) {
    return &User{id, authToken, email, password, name, createdAt, updatedAt}, nil
}


