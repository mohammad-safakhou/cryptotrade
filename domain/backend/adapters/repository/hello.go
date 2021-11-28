package repository

import (
	"cryptotrade/domain/backend/core/ports"
	"database/sql"
)

type helloRepository struct {
	db *sql.DB
}

func NewHelloRepository(db *sql.DB) ports.HelloRepository {
	return helloRepository{db: db}
}

func (hp helloRepository) Get() string {
	return ""
}

func (hp helloRepository) Save(input string) {
	// save
}
