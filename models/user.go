package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserFormat struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) hashedPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := u.hashedPassword()
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Format() UserFormat {
	return UserFormat{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
