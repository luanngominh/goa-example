package model

import (
	uuid "github.com/satori/go.uuid"
)

// IsZero check uuid is zero
func UUIDIsZero(u *uuid.UUID) bool {
	if u == nil {
		return true
	}
	for _, c := range u {
		if c != 0 {
			return false
		}
	}
	return true
}