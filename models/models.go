package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	FirtsName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
}

type Cars struct {
	gorm.Model
	ID     uuid.UUID `json:"id"`
	Branch string    `json:"branch"`
	Year   int       `json:"year"`
	UserId Users     `json:"user_id"`
}
