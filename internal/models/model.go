package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

type Response struct {
	Message     string `json:"message,omitempty"`
	ID          uint   `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
