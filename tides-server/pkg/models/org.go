package models

import (
	"gorm.io/gorm"
)

// User schema
type Org struct {
	gorm.Model

	// city
	City string `json:"city,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// priority
	// Enum: [Low Medium High]
	Priority string `json:"priority,omitempty"`

	// username
	Username string `gorm:"unique;not null"`
}


