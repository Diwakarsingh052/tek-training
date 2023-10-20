package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Inventory struct {
	gorm.Model
	ItemName    string  `json:"item_name"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	UserId      uint    `json:"user_id"`
	CostPerItem float64 `json:"cost_per_item"`
}

// NewInventory contains information needed to create a ShirtInventory.
type NewInventory struct {
	ItemName    string  `json:"item_name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required,number"`
	CostPerItem float64 `json:"cost_per_item" validate:"required,number,gt=0"`
	Category    string  `json:"category" validate:"required"`
}
