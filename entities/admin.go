package entities

import "time"

type Admin struct {
	ID       string    `gorm:"primaryKey;type:varchar(64);"`
	Items    []Item    `gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email    string    `gorm:"type:varchar(128);unique;not null;"`
	Name     string    `gorm:"type:varchar(256);not null;default:'';"`
	Avatar   string    `gorm:"type:varchar(256);not null;default:'';"`
	CreateAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdateAt time.Time `gorm:"not null;autoUpdateTime;"`
}
