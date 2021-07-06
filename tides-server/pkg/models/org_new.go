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


type OrgNew struct {
	gorm.Model

	// org_id
	OrgID uint `gorm:"primary_key" json:"orgID,omitempty"`

	// orgname
	Orgname string `json:"orgname,omitempty"`

	// own_res_id
	OwnResID  pq.Int64Array `json:"ownResID,omitempty"`

	// ban_res_id
	BanResID  pq.Int64Array `json:"banResID,omitempty"`

	// templates applied by this org
	ApplyTempID  pq.Int64Array `json:"applyTempID,omitempty"`
}