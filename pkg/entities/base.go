package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	UUID      uuid.UUID `gorm:"type:uuid;primary_key" json:"uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (g *Base) BeforeCreate(tx *gorm.DB) (err error) {
	g.UUID = uuid.New()
	return nil
}
