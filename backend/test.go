package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"first_name"`
}
