package entities

import "time"

type Playercoin struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID string    `gorm:"type:varchar(64);not null;"`
	Amount   int64     `gorm:"not null;"`
	CreateAt time.Time `gorm:"not null;autoCreateTime;"`
}
