package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Model struct {
	ID uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}