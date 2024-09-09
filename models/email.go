package models

import "gorm.io/gorm"

type Email struct {
	gorm.Model // Embedding gorm.Model to inherit fields like ID, CreatedAt, UpdatedAt, DeletedAt

	Subject string `gorm:"size:250; not null;" json:"subject"`
	Message string `gorm:"size:500; not null" json:"message"`
	Link    string `gorm:"size:250" json:"link"`
}
