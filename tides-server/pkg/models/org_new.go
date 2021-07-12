package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)


type OrgNew struct {
	gorm.Model

	// orgname
	OrgName string `json:"orgName,omitempty" gorm:"uniqueIndex"`

	// own_res_id
	OwnResID  pq.Int64Array `gorm:"type:integer[]"`

	// ban_res_id
	BanResID  pq.Int64Array `gorm:"type:integer[]"`

	// templates applied by this org
	ApplyTempID  pq.Int64Array `gorm:"type:integer[]"`
}