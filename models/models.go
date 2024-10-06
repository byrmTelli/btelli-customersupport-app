package models

import (
	"time"

	"gorm.io/gorm"
)

// Database Models Defined Here.

type User struct {
	ID           uint   `gorm:"primaryKey"`
	UserName     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Name         string `gorm:"not null"`
	Phone        string `gorm:"unique;not null"`
	Surname      string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	RoleID       uint
	Role         Role `gorm:"foreignKey:RoleID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Complaint struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	CategoryID  uint
	Category    ComplaintCategory `gorm:"foreignKey:CategoryID"`
	Status      ComplaintStatus   `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type ComplaintCategory struct {
	ID          uint   `gorm:"primarKey"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Comment struct {
	ID          uint `gorm:"primarKey"`
	ComplaintID uint
	Complaint   Complaint `gorm:"foreignKey:ComplaintID"`
	UserID      uint
	User        User   `gorm:"foreignKey:UserID"`
	CommentText string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
