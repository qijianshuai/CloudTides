package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
	"github.com/lib/pq"
)


type TemplateNew struct {
	gorm.Model

	// templateID
	TemplateID uint `gorm:"primary_key" json:"templateID,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// provider orgID
	ProvideOrgID uint `json:"provideOrgID,omitempty"`

	// VM template ID (VMTemp in previous template.go)
	VmTemplateID pq.Int64Array `json:"vmTemplateID,omitempty"`

	// type
	Type string `json:"type,omitempty"`


}