package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
)


type CloudProviderNew struct {
	gorm.Model

	// cloudProvider_id
	CloudProviderID uint `gorm:"primary_key" json:"cloudPrividerID,omitempty"`

	// org_id
	OrgID uint `json:"orgID,omitempty"`
	
	// org
	Org Org `gorm:"foreignKey:OrgID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// cp_backing_type
	CpBackingType string `json:"cpBackingType,omitempty"`

}