package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID string `json:"id,omitempty" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
	ProviderId string `json:"provider_id" gorethink:"provider_id"`
}

func (u *User) GenerateUUID() {
	u.ID = uuid.New().String()
}