package models

import (
	"github.com/google/uuid"
)

type Provider struct {
	ID string `json:"id,omitempty" gorethink:"id,omitempty"`
	CompanyName string `json:"company_name" gorethink:"company_name"`

}

func (p *Provider) GenerateUUID() {
	p.ID = uuid.New().String()
}