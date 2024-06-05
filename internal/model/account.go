package model

import (
	"crypto/sha256"
	"fmt"
)

type Account struct {
	ModelBase
	Email          string `gorm:"unique" json:"email"`
	Password       string `gorm:"-" json:"password"`
	HashedPassword string `json:"-"`
}

func (a *Account) GenerateHashPassword() string {
	hashStr := fmt.Sprintf("%s:%s", a.Email, a.Password)
	return string(sha256.New().Sum([]byte(hashStr)))
}
