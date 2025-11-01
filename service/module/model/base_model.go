package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:ID" json:"id" form:""`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)" json:"-"`
}
