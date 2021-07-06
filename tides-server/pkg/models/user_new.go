package models

import (
    "encoding/json"

    "github.com/go-openapi/errors"
    "github.com/go-openapi/strfmt"
    "github.com/go-openapi/swag"
    "github.com/go-openapi/validate"
    "gorm.io/gorm"
)


type UserNew struct {
    gorm.Model

    // user_id
    UserID uint `gorm:"primary_key" json:"userID,omitempty"`

    // username
    Username string `json:"username,omitempty"`

    // password
    Password string `json:"password,omitempty"`

    // org_id
    OrgID uint `json:"orgID,omitempty"`
    
    // org
    Org Org `gorm:"foreignKey:OrgID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

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

