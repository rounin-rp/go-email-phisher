package models

import "gorm.io/gorm"

type Links struct {
	gorm.Model // Embedding gorm.Model to inherit fields like ID, CreatedAt, UpdatedAt, DeletedAt

	MagicLink    string `gorm:"size:36; not null; unique"`
	UserId       string `gorm:"size:7; not null" json:"user_id"`
	Email        string `gorm:"size:20; not null" json:"email"`
	HasOpened    bool   `gorm:"not null" json:"has_opened"`
	TimesClicked int    `gorm:"default:0; not null" json:"times_clicked"`
}
