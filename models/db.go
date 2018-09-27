package models

import "time"

type Db struct {
	Id        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	DBSchema  string    `gorm:"column:dbschema"`
	Host      string    `gorm:"column:host"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Port      string    `gorm:"column:port"`
	UserId    int       `gorm:"column:userId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Dbs []Db
