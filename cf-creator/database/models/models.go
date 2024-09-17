package models

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Password string   `json:"password"`
	Role     string   `json:"role"`
	Editors  []Editor `gorm:"foreignKey:OwnerID" json:"editors"`
}

type Editor struct {
	gorm.Model
	OwnerID  uint   `json:"owner_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
