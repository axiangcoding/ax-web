package schema

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	//逻辑主键自增字段
	ID        int64          `gorm:"primary_key;auto_increment:true;bigserial;unique;not null;index" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
