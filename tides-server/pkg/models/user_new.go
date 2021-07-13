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

    // org_name
    OrgName string `gorm:"index" json:"orgName,omitempty"`


    // role SITE_ADMIN/ ORG_ADMIN/ USER
    Role string `json:"role,omitempty"`

    // email
    Email string `json:"email,omitempty"`

    // pw_reset
    PwReset bool `json:"pwReset,omitempty"`

    // phone
    Phone string `json:"phone,omitempty"`

    // avatar //TODO: NOT IMPLEMENTED
    Avatar string `json:"avatar,omitempty"`
}

