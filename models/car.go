package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Car represents the model for an car
type Car struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Car) BeforeCreate(tx *gorm.DB) (err error) {
	// hooks
	fmt.Println("Car Before Create()")
	if len(p.Name) < 4 {
		err = errors.New("car name is too short")
	}

	return
}
