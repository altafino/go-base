package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"

	"github.com/go-pg/pg/orm"
)

// Profile holds specific application settings linked to an Account.
type Profile struct {
	ID        int       `json:"id,omitempty"`
	AccountID int       `json:"-"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Profile) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
		p.UpdatedAt = now
	}
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Profile) BeforeUpdate(db orm.DB) error {
	if err := p.Validate(); err != nil {
		return err
	}
	p.UpdatedAt = time.Now()
	return nil
}

// Validate validates Profile struct and returns validation errors.
func (p *Profile) Validate() error {

	return validation.ValidateStruct(p,
		validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	)
}
