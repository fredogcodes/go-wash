package store

import "gorm.io/gorm"

type Store interface {
	// Add methods here
}

type SQLStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &SQLStore{
		db: db,
	}
}
