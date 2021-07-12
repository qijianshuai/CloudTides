package models

import (
    "gorm.io/gorm"
)


type UserNew struct {
    gorm.Model

    // username
    Username string `gorm:"uniqueIndex" json:"username,omitempty"`

    // password
    Password string `json:"password,omitempty"`

    // org_id
    OrgID uint `gorm:"index" json:"orgID,omitempty" sql:"type:uint REFERENCES OrgNew(OrgID)"`
    
    // org
    Org OrgNew `gorm:"ForeignKey:OrgID;AssociationForeignKey:OrgID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

    // role
    Role string `json:"role,omitempty"`

    // email
    Email string `json:"email,omitempty"`

    // pw_reset
    PwReset bool `json:"pwReset,omitempty"`

    // phone
    Phone string `json:"phone,omitempty"`

    // avatar
    Avatar string `json:"avatar,omitempty"`
}

