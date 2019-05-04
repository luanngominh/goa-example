package model

import (
	uuid "github.com/satori/go.uuid"
)

type Note struct {
	Model

	UserID uuid.UUID `json:"user_id"`
	Data string `json:"data"`
}