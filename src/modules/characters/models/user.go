package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
)

type Character struct {
	ID           int       `json:"id" db:"id"`
	Name    		 string    `json:"first_name" db:"first_name"`
	Description  string    `json:"last_name" db:"first_name"`

	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func(character Character) TableName() string {
	return "characters"
}
