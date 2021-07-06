package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
)


type ResVapp struct {
	gorm.Model

	// vapp_id
	VappID uint `json:"vappID,omitempty"`

	// template_id
	TemplateID uint `json:"templateID,omitempty"`

	// Res_id
	ResID uint `json:"resID,omitempty"`

}