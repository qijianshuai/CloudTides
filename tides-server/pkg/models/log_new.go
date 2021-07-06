package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
	"time"
)


type LogNew struct {
	gorm.Model

	// log_id
	LogID uint `gorm:"primary_key" json:"logID,omitempty"`

	// user_id
	UserID uint `json:"userID,omitempty"`
	
	// operation
	Operation string `json:"operation,omitempty"`

	// Time
	Time time.Time `json:"time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

}