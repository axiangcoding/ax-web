package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId uint `gorm:"uniqueIndex"`
	//下面三种信息在库中都是不可重复的
	UserName string  `gorm:"uniqueIndex;size:255"`
	Email    *string `gorm:"uniqueIndex;size:255"`
	Phone    *string `gorm:"uniqueIndex;size:255"`
	Password string
}
