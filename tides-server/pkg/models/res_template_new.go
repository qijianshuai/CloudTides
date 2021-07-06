package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
)


type ResTemplate struct {
	gorm.Model

	// Res_id
	ResID uint `json:"resID,omitempty"`

	// template_id
	TemplateID uint `json:"templateID,omitempty"`

	// Time
	Time time.Time `json:"time,omitempty"`

	

}