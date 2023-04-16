package models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model         // adds ID, created_at etc.
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}
