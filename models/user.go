package models

import "time"

type User struct {
	Id        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Status    string    `gorm:"column:status"`
	Username  string    `gorm:column:username"`
	Password  string    `gorm:"column:password"`
	Roles     string    `gorm:"column:roles"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Users []User
