package models

import (
	"gorm.io/gorm"
	"time"
)


type ResTemplate struct {
	gorm.Model

	// Res_id
	ResID uint `json:"resID,omitempty"`

	// template_id
	TemplateID uint `json:"templateID,omitempty"`

	// Time
	Time time.Time `json:"time,omitempty"`

	

}