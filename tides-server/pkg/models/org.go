package models

import (
	"gorm.io/gorm"
)

// User schema
type Org struct {
	gorm.Model

	orgName string `gorm:"unique"`
	ownResourceId string
	banResourceId string
	template string
}


