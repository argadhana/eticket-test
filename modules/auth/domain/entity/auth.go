package entity

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:100;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         string    `gorm:"size:50;not null" json:"role"`
	FullName     string    `gorm:"size:150" json:"fullname"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName override biar sesuai dengan nama tabel
func (User) TableName() string {
	return "users"
}