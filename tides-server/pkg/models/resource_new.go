package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gorm.io/gorm"
)


type ResourceNew struct {
	gorm.Model

	// res_id
	ResID uint `gorm:"primary_key" json:"resID,omitempty"`

	// org_id
	OrgID uint `json:"orgID,omitempty"`

	// res_path
	ResPath string `json:"resPath,omitempty"`

	// res_status
	ResStatus string `json:"resStatus,omitempty"`

	// cloudprovider_id
	CloudProviderID uint `json:"cloudProviderID,omitempty"`

	// cloudProvider
	CloudProvider CloudProvider `gorm:"foreignKey:CloudProviderID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}