package models

import (
	"gorm.io/gorm"
)


type CloudProviderNew struct {
	gorm.Model

	// cloudProvider_id
	CloudProviderID uint `gorm:"uniqueIndex" json:"cloudPrividerID,omitempty"`

	// org_id
	OrgID uint `json:"orgID,omitempty" sql:"type:uint REFERENCES OrgNew(ID)"`

	// org
	Org OrgNew `gorm:"ForeignKey:OrgID;AssociationForeignKey:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// cp_backing_type
	CpBackingType string `json:"cpBackingType,omitempty"`

}